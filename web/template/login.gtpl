<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/login" method="post">
            Username:<input type="text" name="username">
            Password:<input type="password" name="password">
            Age:<input type="text" name="age">
            <input type="hidden" name="token" value="{{.}}">
            <input type="submit" value="Login">
        </form>
        <select name="fruit">
        <option value="apple">apple</option>
        <option value="pear">pear</option>
        <option value="banana">banana</option>
        </select>
    </body>
</html>