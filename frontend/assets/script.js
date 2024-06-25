// slider date

const rangeInputCrea = document.querySelectorAll(".range-input input"),
    priceInputCrea = document.querySelectorAll(".price-input input");

const rangeInputAlbum = document.querySelectorAll(".range-input-album input"),
    priceInputAlbum = document.querySelectorAll(".price-input-album input");

let priceGap = 1;

priceInputCrea.forEach(input =>{
    input.addEventListener("input", e =>{
        let minPrice = parseInt(priceInputCrea[0].value),
            maxPrice = parseInt(priceInputCrea[1].value);

        if((maxPrice - minPrice >= priceGap) && maxPrice <= rangeInputCrea[1].max){
            if(e.target.className === "input-min"){
                rangeInputCrea[0].value = minPrice;
            }else{
                rangeInputCrea[1].value = maxPrice;
            }
        }
    });
});

priceInputAlbum.forEach(input =>{
    input.addEventListener("input", e =>{
        let minPrice = parseInt(priceInputAlbum[0].value),
            maxPrice = parseInt(priceInputAlbum[1].value);

        if((maxPrice - minPrice >= priceGap) && maxPrice <= rangeInputCrea[1].max){
            if(e.target.className === "input-min"){
                rangeInputAlbum[0].value = minPrice;
            }else{
                rangeInputAlbum[1].value = maxPrice;
            }
        }
    });
});

rangeInputCrea.forEach(input =>{
    input.addEventListener("input", e =>{
        let minVal = parseInt(rangeInputCrea[0].value),
            maxVal = parseInt(rangeInputCrea[1].value);

        if((maxVal - minVal) < priceGap){
            if(e.target.className === "range-min"){
                rangeInputCrea[0].value = maxVal - priceGap
            } else {
                rangeInputCrea[1].value = minVal + priceGap;
            }
        } else {
            priceInputCrea[0].value = minVal;
            priceInputCrea[1].value = maxVal;
        }
    });
});

rangeInputAlbum.forEach(input =>{
    input.addEventListener("input", e =>{
        let minVal = parseInt(rangeInputAlbum[0].value),
            maxVal = parseInt(rangeInputAlbum[1].value);

        if((maxVal - minVal) < priceGap){
            if(e.target.className === "range-min"){
                rangeInputAlbum[0].value = maxVal - priceGap
            } else {
                rangeInputAlbum[1].value = minVal + priceGap;
            }
        } else {
            priceInputAlbum[0].value = minVal;
            priceInputAlbum[1].value = maxVal;
        }
    });
});

function checkAll() {
    const selectAll = document.getElementById("select-all");
    const checkboxes = document.querySelectorAll('input[type="checkbox"]:not(#select-all)');
    for (let i = 0; i < checkboxes.length; i++) {
        checkboxes[i].checked = selectAll.checked;
    }
}