// app.js - Handles UI and logic for Ping Pong Demo

document.body.style.fontFamily = 'Arial, sans-serif';
document.body.style.textAlign = 'center';
document.body.style.marginTop = '100px';

const h1 = document.createElement('h1');
h1.textContent = 'Ping Pong with Go 🚀';
document.body.appendChild(h1);

const pingBtn = document.createElement('button');
pingBtn.id = 'ping';
pingBtn.textContent = 'PING';
pingBtn.style.backgroundColor = 'green';
pingBtn.style.color = 'white';
pingBtn.style.fontSize = '2rem';
pingBtn.style.padding = '20px 40px';
pingBtn.style.border = 'none';
pingBtn.style.borderRadius = '10px';
pingBtn.style.cursor = 'pointer';
pingBtn.onclick = pingServer;
document.body.appendChild(pingBtn);

document.body.appendChild(document.createElement('br'));

const pongBtn = document.createElement('button');
pongBtn.id = 'pong';
pongBtn.textContent = 'PONG';
pongBtn.style.backgroundColor = 'blue';
pongBtn.style.color = 'white';
pongBtn.style.fontSize = '2rem';
pongBtn.style.padding = '20px 40px';
pongBtn.style.border = 'none';
pongBtn.style.borderRadius = '10px';
pongBtn.style.cursor = 'pointer';
pongBtn.style.marginTop = '20px';
pongBtn.style.display = 'none';
document.body.appendChild(pongBtn);

async function pingServer() {
  let res = await fetch('/ping');
  let text = await res.text();
  console.log('Server says:', text);
  pongBtn.style.display = 'inline-block';
  setTimeout(() => {
    pongBtn.style.display = 'none';
  }, 2000);
}
