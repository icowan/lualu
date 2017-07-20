<!DOCTYPE HTML>
<html>
<head>
    <title>Home</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="keywords" content="" />
    <script type="applijewelleryion/x-javascript"> addEventListener("load", function() { setTimeout(hideURLbar, 0); }, false); function hideURLbar(){ window.scrollTo(0,1); } </script>
    <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel='stylesheet' type='text/css' />
    <!-- Custom Theme files -->
    <link href='http://fonts.useso.com/css?family=Raleway:400,600,700' rel='stylesheet' type='text/css'>
    <link href="/static/css/style.css?v=0.0.01" rel='stylesheet' type='text/css' />
    <!--<link rel="stylesheet" href="https://cdn.bootcss.com/highlight.js/8.0/styles/monokai_sublime.min.css">-->
    <!--<link rel="stylesheet" href="https://lattecake.oss-cn-beijing.aliyuncs.com/static/moderna/css/overwrite.css">-->
    <style>
        video {
            max-width: 100%;
            -ms-interpolation-mode: bicubic;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            width: 100% !important;
            height: 100% !important;
        }
        .container img {
            max-width: 100%;
            height: auto;
        }
    </style>

    {{define "header"}}
    {{end}}
</head>
<body>

{{template "./default_navigation.tpl"}}

{{.LayoutContent}}

{{template "./default_footer.tpl"}}


<div class="copyright wow fadeInDown"  data-wow-duration=".8s" data-wow-delay=".2s">
    <div class="container">
        <p>© LatteCake 2017 All right reserved. By <a href="#">dudulu.me</a></p>
    </div>
</div>

<!--<sciprt src="https://cdn.bootcss.com/jquery/3.1.1/jquery.min.js"></sciprt>-->
<script src="/static/js/jquery-1.11.1.min.js"></script>
<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
<!-- animation-effect -->
<link href="/static/css/animate.min.css" rel="stylesheet">
<script src="/static/js/wow.min.js"></script>
<script>
    new WOW().init();
</script>
<!-- //animation-effect -->
</body>
</html>