Each muninn service has its own direct rabbitmq queueu, so:

Scavenger has muninn.scavenger.queue
Incubator has muninn.incubator.queue
Observer has muninn.observer.queue

If the observer needs to public a message to the scavenger, it does so using the muninn.scavenger.queue
