$(function () {
  // URL解析
  $.getUrlParam = function (name) {
    var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)')
    var r = window.location.search.substr(1).match(reg)
    if (r != null) return unescape(r[2])
    return null
  }

  var userid = $.getUrlParam('userid')
  var pageid = parseInt($.getUrlParam('pageid'))

  if (userid == null) {
    window.location.href = '/'
    return
  }

  $('#before').on('click', () => {
    window.location.href = `/user?userid=${userid}&pageid=${pageid - 1}`
  })

  $('#next').on('click', () => {
    window.location.href = `/user?userid=${userid}&pageid=${pageid + 1}`
  })

  console.log(pageid)

  $('.article').load('/src/card/usercard.html', () => {
    // 发起请求
    infoURL = 'http://127.0.0.1:8080/getarts'

    $('#pageid').text(`第${pageid}页`)

    $.ajax(infoURL, {
      method: 'POST',
      data: {
        userid: userid,
        pageid: pageid,
      },
      // 得到响应的处理
      success: (data) => {
        if (data.code != '200') {
          window.location.href = `/user?userid=${userid}&pageid=1`
          return
        }
        $('#headimg').attr('src', data.data.user.headimg)
        $('#username').text(data.data.user.name)
        $.each(data.data.articles, (index, item) => {
          username = item.username
          title = item.title
          subtitle = item.subtitle
          arttype = item.arttype
          headimg = item.headimg
          likes = item.likes
          share = item.share
          issuper = item.super
          $(`#${index} .username`).text(username)
          $(`#${index} .card-title`).text(title)
          $(`#${index} .subtitle`).html('&emsp;&emsp;' + subtitle)
          $(`#${index} .arttype`).text(arttype)
          $(`#${index} img`).attr('src', headimg)
          $(`#${index} .likes`).text(likes)
          $(`#${index} .share`).text(share)
          $(`#${index}`).on('click', () => {
            window.location.href = `/article?artid=${item.artid}`
          })
          if (issuper == true) {
            $(`#${index} .badage-art`).text('精品')
            $(`#${index} .badage-art`).removeClass('badge-primary')
            $(`#${index} .badage-art`).addClass('badge-success')
          } else {
            $(`#${index} .badage-art`).text('推荐')
            $(`#${index} .badage-art`).removeClass('badge-success')
            $(`#${index} .badage-art`).addClass('badge-primary')
          }
        })
      },
      error: () => {
        $('#cards').html(
          '<h4 class="alert alert-danger m-3 text-center lead">服务器连接失败，可能正在维护，刷新重试或请联系管理员...</h4>'
        )
      },
    })
  })
})
