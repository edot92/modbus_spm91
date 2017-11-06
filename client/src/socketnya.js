// import io from 'socket.io-client';
//
window.socket = window.io('http://localhost');
window.socket.on('connect', function () {
  console.log('sokcet konek');
});
window.socket.on('disconnect', function () {
  console.log('sokcet disconnect');
});
