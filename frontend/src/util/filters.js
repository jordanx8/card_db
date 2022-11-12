export function currentPlayersFilter(player) {
    return player.seasonsPlayed.includes('present')
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