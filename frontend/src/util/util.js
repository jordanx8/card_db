export function getImage(firstName, lastName) {
    let num = "01"
    const twonames = ["Trey_Murphy", "Larry_Nance", "Cameron_Thomas", "Anthony_Davis"]
    for (let i = 0; i < twonames.length; i++) {
        if (firstName + "_" + lastName.replace("-", "_") === twonames[i]) {
            num = "02"
        }
    }
    if (firstName.includes(".")) {
        firstName = firstName.replace(".", "")
    }
    return "https://www.basketball-reference.com/req/202106291/images/players/" + lastName.toLowerCase().slice(0, 5) + firstName.toLowerCase().slice(0, 2) + num + ".jpg"
}

export function nameToFilter(name) {
    return function filterCardsByName(card) {
        return card.playerName === name;
    }
}

export function seasonToFilter(season) {
    return function filterCardsBySeason(card) {
        return season.includes(card.season);
    }
}

export function filterPlayerListDownTo(name) {
    return function filterPlayersByName(player) {
        return ((player.firstName + " " + player.lastName) === name)
    }
}