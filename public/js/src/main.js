import $ from 'jquery'
import Notify from 'handy-notification'

$('.register').on('submit', e => {
  e.preventDefault()

  let
    username = $('.r_username').val(),
    email = $('.r_email').val(),
    password = $('.r_password').val(),
    password_again = $('.r_password_again').val()

  if(!username || !email || !password || !password_again){
    Notify({ value: 'Some values are missing!' })
  } else {
    $.ajax({
      url: '/user/register',
      data: {
        username,
        email,
        password,
        password_again
      },
      method: 'POST',
      dataType: 'JSON',
      success: data => {
        let { mssg, success } = data
        console.log(data)
        Notify({
          value: mssg,
          done: () => success ? location.href = '/login' : null
        })
      }
    })
  }
})

$('.login').on('submit', e => {
  e.preventDefault()
  let
    username = $('.l_username').val(),
    password = $('.l_password').val()

  if(!username || !password){
    Notify({ value: 'Some values are missing!' })
  } else {
    $.ajax({
      url: '/user/login',
      data: {
        username,
        password
      },
      method: 'POST',
      dataType: 'JSON',
      success: data => {
        let { mssg, success, user } = data
        console.log(data)
        Notify({
          value: mssg,
          done: () => success ? location.href = `/profile/${user}` : null
        })
      }
    })
  }
})
