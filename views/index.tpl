<!DOCTYPE html>
<html>
<head lang="en">
  <meta charset="UTF-8">
  <title>Login Page | Amaze UI Example</title>
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
    </form>
    <hr>
  </div>
</div>
</body>
</html>
