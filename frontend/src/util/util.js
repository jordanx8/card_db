export function getImage(firstName, lastName) {
    let num = "01"
    const twonames = ["Trey_Murphy", "Larry_Nance", "Cameron_Thomas", "Anthony_Davis"]
    for (let i = 0; i < twonames.length; i++) {
        if (firstName + "_" + lastName.replace("-", "_") === twonames[i]) {
            num = "02"
        }
    }
    return "https://www.basketball-reference.com/req/202106291/images/players/" + lastName.toLowerCase().slice(0, 5) + firstName.toLowerCase().slice(0, 2) + num + ".jpg"
}