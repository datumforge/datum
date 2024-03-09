var issuer = $("#name-text");
var accountName = $("#email-text");
var submitDets = $("#submit-details");
var btnCont = $("#btn-key__generator");
var totpCont = $("#totp-generate__container");
var detailsCont = $(".details-container");
var keyDetails = $(".input-for__key");
var haveKeyDets = $("#key-input");

totpCont.hide();
detailsCont.hide();
keyDetails.hide();

function showKeyInput() {
  keyDetails.show();
}

function displayDetails() {
  detailsCont.show();
  btnCont.hide();
}

function submitDetails() {
  $.ajax({
    method: "POST",
    url: "/",
    data: {
      data_action: "GENERATE KEY",
      issuer: issuer.val(),
      accountName: accountName.val(),
    },
    success: function () {
      console.log("issuer: ", issuer.val());
      console.log("accountName: ", accountName.val());

      issuer.css("display", "none");
      accountName.css("display", "none");
      submitDets.css("display", "none");
      btnCont.hide();
    },
  });
}

function showPrint() {
  exportBtn.show();
}

function printQR() {
  var dataCont = document.getElementById("printDV");
  dataCont.style.width = "100%";
  dataCont.style.height = "100%";

  // Paper and able size
  var opt = {
    margin: 0.5,
    filename: "QR_Code.pdf",
    image: { type: "jpeg", quality: 1 },
    html2canvas: { scale: 1 },
    jsPDF: {
      unit: "in",
      format: "legal",
      orientation: "portrait",
      precision: "12",
    },
  };

  // Choose the timeManagement and pass it to html2pdf() function and call the save() on it to save as pdf
  html2pdf().set(opt).from(dataCont).save();
}

function validateOTP() {
  // Send a POST request to the server to validate the OTP
  console.log("test", $("#otp-input").text() == $("#totp").val());
  if ($("#otp-input").text() == $("#totp").val()) {
    $("#otp-status").text("TOTP code is valid!");
  } else {
    $("#otp-status").text("Invalid TOTP code!");
  }
}

function updateTimer() {
  var now = new Date();
  var timeLeft = 30 - (now.getSeconds() % 30);
  $("#timer").text("Code expires in \n" + timeLeft);

  if (now.getSeconds() % 30 === 0) {
    // Wait 1 second and reload the page
    setTimeout(function () {
      location.reload();
    }, 5);
  }
}
setInterval(updateTimer, 1000);
