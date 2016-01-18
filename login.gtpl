<html>
  <head>
    <title>Login</title>
  </head>
  <body>
    <form action="/login" method="post">
      Username:<input type="text" name="username">
      Password:<input type="password" name="password">
      Email:<input type="text" name="email">
      </br>
      <input type="radio" name="gender" value="1">Male
      <input type="radio" name="gender" value="2">Female
      </br>
      <select name="sports-team">
        <option value="Giants">Giants</option>
        <option value="As">As</option>
        <option value="49ers">49ers</option>
        <option value="Raiders">Raiders</option>
      </br>
      <input type="submit" value="Login">
    </form>
  </body>
</html>
