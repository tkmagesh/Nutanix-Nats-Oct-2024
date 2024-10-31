<!-- --------------------- -->
<!-- Start 1st cluster -->
nats-server -c india.conf

<!-- Start 2nd cluster -->
nats-server -c us.conf

<!-- --------------------- -->
<!-- Client operations -->

<!-- generate & save the client conf files  -->

nats context save india --server "nats://localhost:4222"

nats context save us --server "nats://localhost:4223"

<!-- Use the generated conf to communicate to the nodes -->

nats --context india reply 'greet' 'hello from india'

nats --context us request 'greet' ''

<!--*** To access the server list, create the client conf with user credentials -->

nats context save india-sys --server "nats://localhost:4222" --user sys --password sys

<!-- Use the india-sys conf to explore the servers -->
nats --context india-sys server list
