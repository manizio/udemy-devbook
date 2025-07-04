$('#login').on('submit', login)

function login(e) {
    e.preventDefault()

    $.ajax({
        url: '/login',
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        window.location = "/home"
    }).fail(function() {
        Swal.fire("Ops...", "Usuário ou senha incorretos", "error")
    })
}
