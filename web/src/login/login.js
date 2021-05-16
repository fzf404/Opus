$(function () {
  // 常用变量
  const loginURL = 'http://127.0.0.1:8080/login'
  const registerURL = 'http://127.0.0.1:8080/register'

  // 获取登录表单
  let $login = $('#login')
  let $register = $('#register')

  // 获取登录消息p标签
  var $lgAlert = $('#login-alert')
  var $rgAlert = $('#register-alert')

  // 处理登录注册
  $login.on('submit', (event) => {
    // 保护不刷新
    event.preventDefault()

    // 获取表单内容
    var name = $("#login [name='name']").val()
    var password = $("#login [name='password']").val()

    // console.log(name, password)

    // 发送登录请求
    $.ajax(loginURL, {
      method: 'POST',
      data: {
        name: name,
        password: password,
      },
      // 得到响应的处理
      success: (data) => {
        $lgAlert.text(data.msg)
        if (data.code != '200') {
          $lgAlert.removeClass()
          $lgAlert.addClass('alert alert-danger mt-3')
        } else {
          $lgAlert.removeClass()
          $lgAlert.addClass('alert alert-success mt-3')
          //存储Token
          window.localStorage.setItem('token', data.data.token)
          // console.log(window.sessionStorage.token)
          window.location.href = '/'
        }
      },
      // 未响应的处理
      error: () => {
        $lgAlert.text('请求处理失败，请联系管理员')
        $lgAlert.removeClass()
        $lgAlert.addClass('alert alert-danger mt-3')
      },
    })
  })

  // 注册处理
  $register.on('submit', (event) => {
    // 保护不刷新
    event.preventDefault()

    // 获取表单内容
    var name = $("#register [name='name']").val()
    var email = $("#register [name='email']").val()
    var password = $("#register [name='password']").val()
    var repassword = $("#register [name='repassword']").val()
    var checkbox = $("#register [value='protocol']")

    // console.log(name, password, repassword, email)

    // 判断两次密码是否一致
    if (!checkbox.is(':checked')) {
      $rgAlert.removeClass()
      $rgAlert.addClass('alert alert-warning mt-3')
      $rgAlert.text('请勾选用户协议')
      return
    }
    if (password != repassword || password == '') {
      $rgAlert.removeClass()
      $rgAlert.addClass('alert alert-warning mt-3')
      $rgAlert.text('两次密码输入不一致')
      return
    }

    // 发送注册请求
    $.ajax(registerURL, {
      method: 'POST',
      data: {
        name: name,
        email: email,
        password: password,
      },
      // 得到响应的处理
      success: (data) => {
        $rgAlert.text(data.msg)
        if (data.code != '200') {
          $rgAlert.removeClass()
          $rgAlert.addClass('alert alert-warning mt-3')
        } else {
          $rgAlert.removeClass()
          $rgAlert.addClass('alert alert-success mt-3')
        }
      },
      // 未响应的处理
      error: () => {
        $rgAlert.text('请求处理失败，请联系管理员')
        $lgAlert.removeClass()
        $lgAlert.addClass('alert alert-danger mt-3')
      },
    })
  })
})
