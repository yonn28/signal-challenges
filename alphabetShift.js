function alphabeticShift(inputString) {
    let splited = inputString.split("");
    let stacked = splited.map((character) => {
        initialCharacter = character.charCodeAt(0)
        initialCharacter = initialCharacter >= 122 ? 96 : initialCharacter;
        return String.fromCharCode(initialCharacter+1);
    })
    return stacked.join("");
}

console.log(alphabeticShift("crazy"))