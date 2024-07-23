variable "master_sg_ids" {
  type        = list(string)
  description = "The security group IDs to be applied to the master nodes."
}

variable "private_network_id" {
  type = string
}

variable "nodes_default_port" {
  type = object({
    network_id = string
    fixed_ips = list(object({
      subnet_id  = string
      ip_address = string
    }))
  })
}

variable "master_port_ids" {
  type = list(string)
}

variable "openstack_disable_floating_ip" {
  type = bool
}

variable "openstack_disable_ingress_floating_ip" {
  type = bool
}