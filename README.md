this is a simple BMI service made in Go language

here are the basic prerequisite to implement/install this service in linux (tested in Ubuntu 20.04.2):
- [Go 1.17.1 package IDE](https://go.dev/dl/)
- [Git package tool](https://git-scm.com/download/linux)
- [systemd/sysvinit daemon](https://man7.org/linux/man-pages/man1/systemd.1.html)

to create the service from scratch please refer to deployment.sh (run as root)
to update the service continuously (such with CI/CD) please refer to integration.sh (run as root)
