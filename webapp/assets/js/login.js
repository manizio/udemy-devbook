$('#login').on('submit', login)

function login(e) {
    e.preventDefault()

    $.ajax({
        url: '/login',
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('passwordl').val(),
        }
    }).done(function() {
        window.location= "/home"
    }).fail(function() {
        alert("usuário ou senha inválidos")
    })
}
