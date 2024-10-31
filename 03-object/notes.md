<!-- Create a bucket -->
nats object add my_store

<!-- Add a file -->
nats object put --name notes my_store notes.md

<!-- List the objects -->
nats object ls my_store

<!-- Remove an object -->
nats object del my_store notes

<!-- Remove a store -->
nats object rm my_store

