# Cluster setup:
There are two important components that a cluster would need to enable CRD versioning:
- [Cert manager](https://book.kubebuilder.io/cronjob-tutorial/cert-manager.html) to manage certificates. API and conversion webhook communicate over TLS, hence the need for certs. 
- [Storage version migrator](https://github.com/kubernetes-sigs/kube-storage-version-migrator) to migrate existing CRs to desired storage version defined in the CRD.

## Creation of the API(defining types.go):

Following the quickstart from [kubebuilder](https://book.kubebuilder.io/quick-start.html?highlight=guestboo#create-an-api), scaffolded the files. (you don't need to do this if you're cloning the repo)
- kubebuilder create api --group webapp --version v1 --kind Guestbook


## Code is segmented in branches. 
- In the v1 branch, "webapp.example.com/v1" API will be the stored version and "webapp.example.com/v2" will be introduced to users. Additionally, we could add a deprecation warning on "webapp.example.com/v1" and start requesting users to move to "webapp.example.com/v2".
- In the v2 branch, "webapp.example.com/v2" will be set as the stored version. Users can still view objects in "webapp.example.com/v1" but they should start migrating objects to  "webapp.example.com/v2" using Storage Version Migrator. 
- In the v3, "webapp.example.com/v1" will no longer be a served version, therefore deprecating it. Users can no longer view objects in v1. 
