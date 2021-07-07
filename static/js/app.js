// Example starter JavaScript for disabling form submissions if there are invalid fields
function applyValidateForm() {
    'use strict'

    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    let forms = document.querySelectorAll('.needs-validation')

    // Loop over them and prevent submission
    Array.prototype.slice.call(forms)
        .forEach(function (form) {
            form.addEventListener('submit', function (event) {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                form.classList.add('was-validated')
            }, false)
        })
}


function notify(msg, type) {
    notie.alert({
        type: type, // optional, default = n4, eum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
        text: msg,
    })
}


function notifyModal(title, text, icon, confirmButtonText) {
    Swal.fire({
        title: title,
        text: text,
        icon: icon,
        confirmButtonText: confirmButtonText
    })
}

function attention() {
    let toast = function (c) {
        const {
            msg = "",
            icon = "success",
            position = "top-end"
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            icon: icon,
            position: position,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({
            icon: 'success',
            title: 'Signed in successfully'
        })
    }
    let success = function (c) {
        const {
            title = "",
            msg = "",
            footer = ""
        } = c;
        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer
        })
    }

    let error = function (c) {
        const {
            title = "",
            msg = "",
            footer = ""
        } = c;
        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer
        })
    }

    let custom = async function (c) {
        const {
            title = "",
            msg = "",
        } = c;

        const {value: result} = await Swal.fire({
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () => {
                if (c.willOpen !== undefined){
                    c.willOpen()
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined){
                    c.didOpen()
                }
            },
            preConfirm: () => {
                // return [
                //     document.getElementById('start').value,
                //     document.getElementById('end').value
                // ]
            }
        })

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel){
                if (result.value !== ""){
                    if (c.callback !== undefined){
                        c.callback(result)
                    }
                }else{
                    c.callback(false)
                }
            }else {
                c.callback(false)
            }
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}

// Test
//
// document.getElementById("test_btn").addEventListener("click", function () {
//     let html = `
//         <form id=="check-availability-form" action="" method="post" novalidate class="needs-validation">
//             <div class="row">
//                 <div class="col">
//                     <div class="row" id="reservation-dates-modal">
//                         <div class="col">
//                             <input disabled autocomplete="off" required class="form-control" name="start" id="start" placeholder="Arrival" type="text" >
//                         </div>
//                         <div class="col">
//                             <input disabled autocomplete="off" required class="form-control" name="end" id="end" placeholder="Departure" type="text" >
//                         </div>
//                     </div>
//                 </div>
//             </div>
//         </form>
//     `
//
//
//     attention().custom({msg: html}).then(undefined)
// })