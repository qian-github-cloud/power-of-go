- setting hosts file
192.168.3.27 kubernetes-master
192.168.3.28 kubernetes-work-1
192.168.3.29 kubernetes-work-2


- setting firewall
- master node 


sudo ufw enable

sudo ufw allow 22/tcp && 
ufw allow 6443/tcp &&
ufw allow 2379/tcp &&
sudo ufw allow 2380/tcp &&
sudo ufw allow 10250/tcp &&
sudo ufw allow 10257/tcp &&
sudo ufw allow 10259/tcp &&
sudo ufw reload


- work node
ufw allow 22/tcp && 
sudo ufw allow 10250/tcp &&
sudo ufw allow 30000:32767/tcp &&
sudo ufw reload


- iptables
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF

sudo sysctl --system


- disable swap

sudo swapoff -a
sudo nano /etc/fstab




- install containerd
wget https://github.com/containerd/containerd/releases/download/v1.6.2/containerd-1.6.2-linux-amd64.tar.gz

sudo tar Czxvf /usr/local containerd-1.6.2-linux-amd64.tar.gz

wget https://raw.githubusercontent.com/containerd/containerd/main/containerd.service

sudo mv containerd.service /usr/lib/systemd/system/

sudo systemctl daemon-reload

sudo systemctl enable --now containerd

sudo systemctl status containerd

- intall runC
wget https://github.com/opencontainers/runc/releases/download/v1.1.1/runc.amd64
sudo install -m 755 runc.amd64 /usr/local/sbin/runc

- setting the config.toml of containerd
sudo mkdir -p /etc/containerd/
containerd config default | sudo tee /etc/containerd/config.toml
sudo sed -i 's/SystemdCgroup \= false/SystemdCgroup \= true/g' /etc/containerd/config.toml
sudo systemctl restart containerd


- Setup
sudo apt update
sudo apt install -y apt-transport-https ca-certificates curl
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list


- install kubeadm kubectl
sudo apt update
sudo apt install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl


- kubeadm init

sudo kubeadm init --pod-network-cidr=192.168.16.0/24
sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=all 
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.3.27:6443 --token w2pojj.fxgihnm2qg8itgv1 \
	--discovery-token-ca-cert-hash sha256:3ce15f2aedc0efd617ef186e0dd6e4a7c134d90517da034cf0c5b11da191534a

kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml
- troublshooting
The connection to server <host>:6443 was refused
sudo -i
swapoff -a
exit
strace -eopenat kubectl version k


- instnall calico



- where calico pod is pending

calico https://qiita.com/fastwind/items/b4229a3f360e7cd5b8e2

kubectl describe nodes

...
Taints:             node-role.kubernetes.io/control-plane:NoSchedule

kubectl taint node kubernetes-master node-role.kubernetes.io/control-plane:NoSchedule-

-- install calicoctl tool
curl -L https://github.com/projectcalico/calico/releases/download/v3.23.1/calicoctl-linux-amd64 -o calicoctl
chmod +x ./calicoctl

-- join master 
kubeadm token create --print-join-command
