# Cluster setup:
There are two important components that a cluster would need to enable CRD versioning:
- [Cert manager](https://book.kubebuilder.io/cronjob-tutorial/cert-manager.html) to manage certificates. API and conversion webhook communicate over TLS, hence the need for certs. 
- [Storage version migrator](https://github.com/kubernetes-sigs/kube-storage-version-migrator) to migrate existing CRs to desired storage version defined in the CRD.

## Creation of the API(defining types.go):

Following the quickstart from [kubebuilder](https://book.kubebuilder.io/quick-start.html?highlight=guestboo#create-an-api)
- kubebuilder create api --group webapp --version v1 --kind Guestbook

### TLDR:
- In the v1 branch, "webapp.example.com/v1" API will be the stored version and  "webapp.example.com/v2" will be introduced to users. Additionally, we could add a deprecation warning on "webapp.example.com/v1" and start requesting users to move to "webapp.example.com/v2". 
- In the v2 branch, "webapp.example.com/v2", 