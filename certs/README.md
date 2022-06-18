## Generate Self-Signed Certificate

1. For the controller.  
   Run: `openssl req -nodes -new -x509 -keyout controller.key -out controller.cert`
2. For the agent.  
   Run: `openssl req -nodes -new -x509 -keyout agent.key -out agent.cert`
3. Copy them to projects.

```bash
cp -rv certs controller/certs
cp -rv certs agent/controller/certs
```
