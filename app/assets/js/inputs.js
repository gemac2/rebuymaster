const inputs = {
    setUp: () => {
        inputs.formatNumbersInput()
        inputs.formatUpperCaseInput()
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
}

export default inputs;
inputs.setUp();