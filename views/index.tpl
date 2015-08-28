<!DOCTYPE html>
<html>
<head>
<title>PainCompiler</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<style>
html,body {
  -webkit-font-smoothing: antialiased;
  font: normal 12px/14px sans-serif;
  width: 100%;
  height: 100%;
  margin: 0;
  overflow: hidden;
   color: #fff; 
  
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
body {
  background: black;
}
</style>
<script> 
var paincompiler = {{.hex}}
console.log({{.xeh}})
</script>
<script async="" type="text/javascript" charset="UTF-8" src="/static/js/matrix.js">
</script>
</head>
<body>
<canvas id="canvas"></canvas>
</body>
</html>
