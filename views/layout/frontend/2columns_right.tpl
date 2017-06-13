<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>你的指尖有改变世界的力量 - markbest.site</title>
    <link rel="icon" href="/static/images/favicon.ico" type="image/x-icon"/>
    <link rel="shortcut icon" href="/static/images/favicon.ico" type="image/x-icon"/>
    <link href="/static/css/customer.css" rel="stylesheet">
    <link href="/static/css/app.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <script src="/static/js/jquery-1.9.1.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
</head>
<body>
    {{template "layout/frontend/header.tpl" .}}
    <div class="container">
        <div class="main_container">
            <div class="left_content col-md-8">
                {{.LayoutContent}}
            </div>
            <div class="right_content col-md-4">
                {{template "layout/frontend/sidebar.tpl" .}}
            </div>
        </div>
    </div>
    {{template "layout/frontend/footer.tpl" .}}
</body>
</html>