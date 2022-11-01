variable "vm_name" {
   type = string
   default = "alpine_base"
}

variable "output_dir" {
   type = string
   default = "./iso"
}

locals{
   vm_name = "${var.vm_name}.qcow2"
   image_dir = "${path.root}/images"
}

source "qemu" "alpine" {
   iso_url = "https://dl-cdn.alpinelinux.org/alpine/v3.15/releases/x86/alpine-standard-3.15.4-x86.iso"
   iso_checksum = "sha256:a0c4dcf7d5afb9090fb035910653dc53b4237222f7dd0bbffb8e331d14a37a62"
   output_directory = "${var.output_dir}"
   qemuargs         = [["-m", "2G"], ["-smp", "cpus=2"]]
   shutdown_command = "poweroff"
   disk_size = "5000M"
   format = "qcow2"
   accelerator = "kvm"
   http_directory = "${path.root}/http"
   ssh_username = "root"
   ssh_password = "rootroot"
   ssh_timeout = "5m"
   vm_name = "${local.vm_name}"
   net_device = "virtio-net"
   disk_interface = "virtio"
   headless = "true"
   boot_wait = "10s"
   vnc_bind_address= "0.0.0.0"
   memory = 1024
   boot_command = [
      "<wait1s>root<enter><wait1s>",
      "<wait1s>setup-interfaces<enter><enter><enter><enter><wait1s>",
      "ifup eth0<enter><wait2>",
      "wget http://{{ .HTTPIP }}:{{ .HTTPPort }}/answers<enter><wait1s>",
      "<wait1s>setup-alpine -f answers<enter><wait5>",
      "rootroot<enter><wait1s>",
      "rootroot<enter><wait5>",
      "<wait1s>y<enter><wait60>",
      "rc-service sshd stop<enter>",
      "mount /dev/vda3 /mnt<enter>",
      "echo 'PermitRootLogin yes' >> /mnt/etc/ssh/sshd_config<enter>",
      "reboot<enter>"
   ]
}

build {
   sources = ["source.qemu.alpine"]

   provisioner "file" {
      source = "alpine_base.sh"
      destination = "/tmp/provisioner.sh"
   }

   provisioner "shell" {
      inline = [". /tmp/provisioner.sh"]
   }

   post-processors {
      post-processor "shell-local" {
         inline = [ "mkdir -p ${local.image_dir}","mv ${var.output_dir}/${local.vm_name} ${local.image_dir}/${local.vm_name}", "rm -r ${var.output_dir}"]
      }
   }
}
