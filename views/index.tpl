<!DOCTYPE html>
<html>
<head lang="en">
  <meta charset="UTF-8">
  <title>TinyURL</title>
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="format-detection" content="telephone=no">
  <meta name="renderer" content="webkit">
  <meta http-equiv="Cache-Control" content="no-siteapp" />
  <link rel="alternate icon" type="image/png" href="/static/i/favicon.png">
  <link rel="stylesheet" href="/static/css/amazeui.min.css"/>
  <style>
    .header {
      text-align: center;
    }
    .header h1 {
      font-size: 200%;
      color: #333;
      margin-top: 30px;
    }
    .header p {
      font-size: 14px;
    }
  </style>
</head>
<body>
<div class="header">
  <div class="am-g">
    <h1>TinyURL</h1>
    <p>A URL shortner by JerryLeooo</p>
  </div>
  <hr />
</div>
<div class="am-g">
  <div class="am-u-lg-6 am-u-md-8 am-u-sm-centered">
    <form method="post" class="am-form" action="/">
      <label for="url">原始URL：</label>
      <input type="text" name="url" id="url" value="{{ .URL }}" />
      <label for="short_url">短URL：</label>
      <input type="text" name="short_url" id="short_url" value="{{ .ShortURL }}" />
      <br>
      <div class="am-cf">
        <input type="submit" name="" value="提交" class="am-btn am-btn-primary am-btn-sm am-fl">
      </div>
      {{ if .ShortURL }}
      <hr>
      访问 <a href="http://{{ .ShortURL }}">{{ .ShortURL }}</a> 体验
      {{ end }}
    </form>
    <hr>
    <script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
      (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

      ga('create', 'UA-99327082-1', 'auto');
      ga('send', 'pageview');
    </script>
  </div>
</div>
</body>
</html>
