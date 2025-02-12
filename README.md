# Muninn

Go application to retrieve hardware information of machines.

![](./img/logo_small.png)

# Architecture

Muninn is a micro-server architectured application. The core application consists of the apiserver, incubator and the scavenger. These are the heart of the app and are meant to run in Kubernetes. They communicate / exchange data through RabbitMQ. The agent runs on the machines of which you want to retrieve the hardware; the agent exposes an API which the scavenger uses when prodded to by the apiserver. Muninnctl is a bubbleatea application, which serves as the main client to use Muninn with. Muninnctl makes api calls to the apiserver, depending on the action the apiserver forwards certain calls to the incubator or scavenger to retrieve and store hardware information.
