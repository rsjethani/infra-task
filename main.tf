provider "libvirt" {
    uri = "qemu:///system"
}

module "cloudinit" {
    source  = "rsjethani/cloudinit/libvirt"
    version = "1.0.0"
    iso_name = var.project_name
    public_key = file(var.public_key_path)
}

resource "libvirt_volume" "server" {
    name = var.project_name
    base_volume_name = var.base_image
    size = 5368709120
}

resource "libvirt_domain" "server" {
    cloudinit = module.cloudinit.id
    name = var.project_name
    memory = 1024
    vcpu = 1
    
    network_interface {
	network_name = "default"
	wait_for_lease = true
    }
    
    disk {
	volume_id = libvirt_volume.server.id
    }
    
    console {
	type        = "pty"
	target_port = "0"
	target_type = "serial"
    }

    graphics {
	type        = "vnc"
	listen_type = "address"
    }
}

