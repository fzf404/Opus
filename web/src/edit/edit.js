$(function () {
  // 编辑器初始化
  const E = window.wangEditor
  const editor = new E('#div1')
  // 设置高度
  editor.config.height = 700
  // 模态框覆盖
  editor.config.zIndex = 100
  // 图片上传
  editor.config.uploadImgServer = 'http://127.0.0.1:8080/upimg'
  // 文件大小限制
  editor.config.uploadImgMaxSize = 2 * 1024 * 1024
  // 类型限制
  editor.config.uploadImgAccept = ['jpg', 'jpeg', 'png', 'gif']
  // 最大文件大小限制
  editor.config.uploadImgMaxLength = 1
  // 上传key值
  editor.config.uploadFileName = 'img'
  // token设置
  editor.config.uploadImgHeaders = {
    Authorization: 'Bearer ' + window.localStorage.getItem('token'),
  }

  editor.create()

  // 事件处理
  $('#upimg').on('click', (event) => {
    event.preventDefault()
    $('#uploadimg').trigger('click')
    console.log('图片上传')
  })

  // 上传事件处理
  $('#uploadimg').on('change', (event) => {
    // console.log(event.target.files[0].name);
    var uploadURL = 'http://127.0.0.1:8080/upimg'
    var formData = new FormData()
    formData.append('img', event.target.files[0])
    // ajax上传图片
    $.ajax(uploadURL, {
      method: 'POST',
      data: formData,
      contentType: false,
      processData: false,
      beforeSend: (xhr) => {
        token = window.localStorage.getItem('token')
        xhr.setRequestHeader('Authorization', 'Bearer ' + token)
      },
      // 得到响应的处理
      success: (data) => {
        if (data.code != '200') {
          alert(data.msg)
        }
        var imgPath = data.data[0]
        $('#imgpath').val(imgPath)
      },
      // 未响应的处理
      error: () => {
        alert('服务器可能正在维护，请刷新重试或联系管理员~')
      },
    })
  })

  // 发布按钮点击事件
  $('#publish').on('click', (event) => {
    addURL = 'http://127.0.0.1:8080/addart'
    var title = $('form [name="title"]').val()
    var subtitle = $('form [name="subtitle"]').val()
    var arttype = $('form [name="arttype"]').val()
    var headimg = $('#imgpath').val()
    var content = editor.txt.html()
    // console.log(title,subtitle,arttype ,headimg, content);
    $.ajax(addURL, {
      method: 'POST',
      data: {
        title: title,
        subtitle: subtitle,
        type: arttype,
        headimg: headimg,
        content: content,
      },
      beforeSend: (xhr) => {
        token = window.localStorage.getItem('token')
        xhr.setRequestHeader('Authorization', 'Bearer ' + token)
      },
      // 得到响应的处理
      success: (data) => {
        alert(data.msg)
        if (data.code == '200') {
          window.location.href = '/'
        }
      },
      error: (data) => {
        alert('服务器可能正在维护，请刷新重试或联系管理员')
      },
    })
  })
})
