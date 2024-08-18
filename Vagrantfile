# Vagrantfile

Vagrant.configure("2") do |config|
  # Choose the base box
  config.vm.box = "ubuntu/focal64"

  # Network configuration (optional, you can customize this)
  config.vm.network "private_network", type: "dhcp"

  # Provisioning with a shell script
  config.vm.provision "shell", path: "provision.sh"
  
  # Forwarding ports for external access (optional)
  config.vm.network "forwarded_port", guest: 80, host: 8080
end
