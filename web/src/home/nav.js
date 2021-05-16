$(function () {
  $('#new').on('click', () => {
    getActive('http://127.0.0.1:8080/active?datacode=0')
  })
  $('#supernew').on('click', () => {
    getActive('http://127.0.0.1:8080/active?datacode=1')
  })
  $('#technologynew').on('click', () => {
    getActive('http://127.0.0.1:8080/active?datacode=2')
  })
  $('#schoolnew').on('click', () => {
    getActive('http://127.0.0.1:8080/active?datacode=3')
  })
  $('#artnew').on('click', () => {
    getActive('http://127.0.0.1:8080/active?datacode=4')
  })
  $('#nonew').on('click', () => {
    getActive('http://127.0.0.1:8080/active?datacode=10')
  })
  $('#testnew').on('click', () => {
    getActive('http://127.0.0.1:8080/active?datacode=404')
    $('#testnew').addClass('btn-success')
  })

  // 处理函数
  function getActive(activeURL) {
    console.log('run..')
    $.ajax(activeURL, {
      method: 'GET',
      // 得到响应的处理
      success: (data) => {
        if (data.code != '200') {
          alert('该分区暂时没有文章哦~')
          return
        }
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
          `<h4 class="alert alert-danger m-3 text-center lead">服务器连接失败，可能正在维护，刷新重试或请联系管理员...</h4>`
        )
      },
    })
  }
})
