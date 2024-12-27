The observer exposes an API endpoint which clients can use to POST ip adress for muninn to scan. So the observer exposes /scan and when a IP has been posted to this endpoint the following flow will happen:

observer sends the ip to the scanvenger -> scavenger tries to connect to the agent on the host -> scavenger sends the json encoded dmidecode to the incubator -> incubator saves the json to the filesystem
