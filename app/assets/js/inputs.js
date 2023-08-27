const inputs = {
    setUp: () => {
        inputs.formatNumbersInput()
        inputs.formatUpperCaseInput()
        inputs.handleEvent()
    },

    formatNumbersInput: () => {
        $(document.body).on('focus', '.only-numbers', function () {
            let decimals = $(this).data("decimals") ? $(this).data("decimals") : 0;
            let padZeros = $(this).data("padzeros") ? true : false;

            var numberMask = IMask(this, {
                mask: Number,
                scale: decimals,
                signed: false,
                thousandsSeparator: '',
                normalizeZeros: false,
                min: 0,
                max: 999999999.99,
                padFractionalZeros: padZeros,
                radix: '.',
            }).on('accept', () => {
                this.innerHTML = numberMask.masked.number;
            });
        })
    },

    formatUpperCaseInput: () => {
        $(document.body).on('input', '.input-uppercase', function () {
            $(this).val($(this).val().toUpperCase());
        })
    },

    handleEvent: () => {
        $('body').on('click', '.form-switch', function () {
            let isChecked = $(this).find(".custom-control-input").is(":checked");

            if (!isChecked) {
                $(`input[type="hidden"][name="IsBuybacksEnabled"]`).prop("value", false)
                $(`input[type="checkbox"][id="flexSwitchCheckDefault"]`).prop("checked", false)
            } else {
                $(`input[type="hidden"][name="IsBuybacksEnabled"]`).prop("value", true)
                $(`input[type="checkbox"][id="flexSwitchCheckDefault"]`).prop("checked", true)
            }
        });

        $('body').on('click', '.position-check', function () {
            let isChecked = $(this).find(".custom-control-input").is(":checked");

            if (!isChecked) {
                $(`input[type="hidden"][name="IsOrderPosition"]`).prop("value", false)
                $(`input[type="checkbox"][id="CheckOrderPosition"]`).prop("checked", false)
            } else {
                $(`input[type="hidden"][name="IsOrderPosition"]`).prop("value", true)
                $(`input[type="checkbox"][id="CheckOrderPosition"]`).prop("checked", true)
            }
        });

        $('body').on('click', '.tp-check', function () {
            let isChecked = $(this).find(".custom-control-input").is(":checked");

            if (!isChecked) {
                $(`input[type="hidden"][name="TakeProfitAchieved"]`).prop("value", false)
                $(`input[type="checkbox"][id="CheckTakeProfit"]`).prop("checked", false)
            } else {
                $(`input[type="hidden"][name="TakeProfitAchieved"]`).prop("value", true)
                $(`input[type="checkbox"][id="CheckTakeProfit"]`).prop("checked", true)
            }
        });

        $('body').on('click', '.sl-check', function () {
            let isChecked = $(this).find(".custom-control-input").is(":checked");

            if (!isChecked) {
                $(`input[type="hidden"][name="StopLossTaken"]`).prop("value", false)
                $(`input[type="checkbox"][id="CheckSl"]`).prop("checked", false)
            } else {
                $(`input[type="hidden"][name="StopLossTaken"]`).prop("value", true)
                $(`input[type="checkbox"][id="CheckSl"]`).prop("checked", true)
            }
        });

        $('body').on('click', '.tw-check', function () {
            let isChecked = $(this).find(".custom-control-input").is(":checked");

            if (!isChecked) {
                $(`input[type="hidden"][name="TradeWon"]`).prop("value", false)
                $(`input[type="checkbox"][id="CheckTradeWon"]`).prop("checked", false)
            } else {
                $(`input[type="hidden"][name="TradeWon"]`).prop("value", true)
                $(`input[type="checkbox"][id="CheckTradeWon"]`).prop("checked", true)
            }
        });

        $('body').on('click', '.tl-check', function () {
            let isChecked = $(this).find(".custom-control-input").is(":checked");

            if (!isChecked) {
                $(`input[type="hidden"][name="TradeLoss"]`).prop("value", false)
                $(`input[type="checkbox"][id="CheckTradeLoss"]`).prop("checked", false)
            } else {
                $(`input[type="hidden"][name="TradeLoss"]`).prop("value", true)
                $(`input[type="checkbox"][id="CheckTradeLoss"]`).prop("checked", true)
            }
        });
    }
}

export default inputs;
inputs.setUp();