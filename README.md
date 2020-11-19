# k6-extension-say

This is a [k6](https://github.com/loadimpact/k6) extension using the
[xk6](https://github.com/k6io/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
|------|

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Clone `xk6`:
  ```shell
  git clone https://github.com/k6io/xk6.git
  cd xk6
  ```

2. Build the binary:
  ```shell
  CGO_ENABLED=1 go run ./cmd/xk6/main.go build master \
    --with github.com/dgzlopes/k6-extension-say
  ```

## Example

```javascript
// script.js
import say from 'k6/x/say';

export default function () {
  console.log(say.hello("Daniel"));
}
```

Result output:

```
$ ./k6 run script.js
          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: ../example.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0000] Hello Daniel                                  source=console

running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU

    data_received........: 0 B 0 B/s
    data_sent............: 0 B 0 B/s
    iteration_duration...: avg=72.55µs min=72.55µs med=72.55µs max=72.55µs p(90)=72.55µs p(95)=72.55µs
    iterations...........: 1   73.974827/s
```
