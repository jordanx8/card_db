import {
  DataTable,
  Table,
  TableHead,
  TableRow,
  TableExpandRow,
  TableExpandedRow,
  TableHeader,
  TableExpandHeader,
  TableBody,
  TableCell,
  TableToolbar,
  TableToolbarSearch,
  TableToolbarContent,
} from '@carbon/react';
import React, {useState} from "react";
import './App.scss';
import { playerHeaders, seasonHeaders, cardHeaders } from './Headers.js';

function App() {
  let playerNames = []
  let playerSeasons = []

  const [json, setJson] = useState(require('./Pelicans.json'))

  function searchJSON(search){
    setJson(require('./Pelicans.json'))
    let searchTerms = search.target.value.split(" ")
    if(searchTerms.length > 1)
    {
      for(let i = 0; i < searchTerms.length; i++){
        setJson(cards => cards.filter(card => JSON.stringify(card).toLowerCase().includes(searchTerms[i].toLowerCase())))
      }
      return
    }
    setJson(cards => cards.filter(card => JSON.stringify(card).toLowerCase().includes(search.target.value.toLowerCase())))
  }

  const playerrows = json.map(getPlayers).filter(checkNil)

  function getPlayers(card) {
    let i = 0;
    while (i < playerNames.length) {
      if (playerNames[i] === (card.FirstName + " " + card.LastName)) {
        return
      }
      i++
    }
    playerNames.push(card.FirstName + " " + card.LastName)
    return {
      id: card.FirstName + " " + card.LastName,
      firstname: card.FirstName,
      lastname: card.LastName,
      seasonsplayed: card.SeasonsPlayed,
      numcards: 0
    }
  }

  function checkNil(card) {
    return (card !== undefined);
  }

  const images = Object.assign({}, ...playerrows.map(getImage))

  function getImage(x) {
    let num = "01"
    const twonames = ["Trey_Murphy", "Larry_Nance", "Cameron_Thomas"]
    for (let i = 0; i < twonames.length; i++) {
      if (x.firstname + "_" + x.lastname.replace("-", "_") === twonames[i]) {
        num = "02"
      }
    }
    return { [x.firstname + "_" + x.lastname.replace("-", "_")]: "https://www.basketball-reference.com/req/202106291/images/players/" + x.lastname.toLowerCase().slice(0, 5) + x.firstname.toLowerCase().slice(0, 2) + num + ".jpg" }
  }

  const seasonobjects = json.map(getSeasons).filter(checkNil)

  function getSeasons(card) {
    let i = 0;
    while (i < playerSeasons.length) {
      if (playerSeasons[i] === (card.FirstName + "_" + card.LastName.replace("-", "_") + "." + card.Season)) {
        return
      }
      i++
    }
    playerSeasons.push(card.FirstName + "_" + card.LastName.replace("-", "_") + "." + card.Season)
    return {
      id: card.FirstName + "_" + card.LastName.replace("-", "_") + "." + card.Season,
      season: card.FirstName + "_" + card.LastName.replace("-", "_") + "." + card.Season
    }
  }

  const seasonrows = Object.assign({}, ...playerrows.map(createSeasonRows))

  function createSeasonRows(playerrow) {
    let temparr = []
    let i = 0
    while (i < playerSeasons.length) {
      if (playerrow.firstname + "_" + playerrow.lastname.replace("-", "_") === playerSeasons[i].split(".")[0]) {
        temparr.push(seasonobjects[i])
      }
      i++
    }
    return { [playerrow.firstname + "_" + playerrow.lastname.replace("-", "_")]: temparr }
  }

  const cardrows = Object.assign({}, ...playerSeasons.map(createCardRows))

  function createCardRows(playerseason) {
    let temparr = []
    let i = 0
    while (i < json.length) {
      if ((json[i].FirstName + "_" + json[i].LastName.replace("-", "_") + "." + json[i].Season) === playerseason) {
        temparr.push({
          id: json[i].FirstName + json[i].LastName.replace("-", "_") + json[i].CardNumber + json[i].Set + json[i].Season + json[i].Insert + json[i].Parallel,
          manufacturer: json[i].Manufacturer,
          set: json[i].Set,
          insert: json[i].Insert,
          parallel: json[i].Parallel,
          number: json[i].CardNumber,
          notes: json[i].Notes
        })
      }
      i++
    }
    return { [playerseason]: temparr }
  }

  return (
    <DataTable rows={playerrows} headers={playerHeaders}>
      {({ rows, headers, getTableProps, getHeaderProps, getRowProps, getBatchActionProps }) => (
        <>
          <TableToolbar>
            <TableToolbarContent>
              <TableToolbarSearch
                tabIndex={getBatchActionProps().shouldShowBatchActions ? -1 : 0}
                onChange={searchJSON}
              />
            </TableToolbarContent>
          </TableToolbar>
          <Table {...getTableProps()}>
            <TableHead>
              <TableRow>
                <TableExpandHeader></TableExpandHeader>
                <TableHeader></TableHeader>
                {headers.map((header) => (
                  <TableHeader  {...getHeaderProps({ header, isSortable: true })}>
                    {header.header}
                  </TableHeader >
                ))}
              </TableRow>
            </TableHead>
            <TableBody>
              {rows.map((row) => (
                <>
                  <TableExpandRow {...getRowProps({ row })}>
                    <TableCell><img src={images[row.cells[0].value + '_' + row.cells[1].value.replace("-", "_")]} alt={row.cells[0].value + ' ' + row.cells[1].value}></img></TableCell>
                    {row.cells.map((cell) => (
                      <TableCell key={cell.id}>{cell.value}</TableCell>
                    ))}
                  </TableExpandRow>
                  {row.isExpanded && (
                    <TableExpandedRow colSpan={headers.length + 2}>
                      <DataTable rows={seasonrows[row.cells[0].value + '_' + row.cells[1].value.replace("-", "_")]} headers={seasonHeaders}>
                        {({ rows, headers, getTableProps, getHeaderProps, getRowProps }) => (
                          <>
                            <Table {...getTableProps()}>
                              <TableHead>
                                <TableRow>
                                  {headers.map((header) => (
                                    <TableHeader colSpan={headers.length + 1}  {...getHeaderProps({ header, isSortable: true })}>
                                      {header.header}
                                    </TableHeader >
                                  ))}
                                </TableRow>
                              </TableHead>
                              <TableBody>
                                {rows.map((row) => (
                                  <>
                                    <TableExpandRow {...getRowProps({ row })}>
                                      {row.cells.map((cell) => (
                                        <TableCell key={cell.id}>{cell.value.split(".")[1]}</TableCell>
                                      ))}
                                    </TableExpandRow>
                                    {row.isExpanded && (
                                      <TableExpandedRow colSpan={headers.length + 2}>
                                        <DataTable rows={cardrows[row.cells[0].value]} headers={cardHeaders}>
                                          {({ rows, headers, getTableProps, getHeaderProps, getRowProps }) => (
                                            <Table {...getTableProps()}>
                                              <TableHead>
                                                <TableRow>
                                                  {headers.map((header) => (
                                                    <TableHeader  {...getHeaderProps({ header, isSortable: true })}>
                                                      {header.header}
                                                    </TableHeader >
                                                  ))}
                                                </TableRow>
                                              </TableHead>
                                              <TableBody>
                                                {rows.map((row) => (
                                                  <TableRow {...getRowProps({ row })}>
                                                    {row.cells.map((cell) => (
                                                      <TableCell key={cell.id}>{cell.value}</TableCell>
                                                    ))}
                                                  </TableRow>
                                                ))}
                                              </TableBody>
                                            </Table>
                                          )}
                                        </DataTable>
                                      </TableExpandedRow>
                                    )}
                                  </>
                                ))}
                              </TableBody>
                            </Table>
                          </>
                        )}
                      </DataTable>
                    </TableExpandedRow>
                  )}
                </>
              ))}
            </TableBody>
          </Table>
        </>
      )}
    </DataTable>
  );
}

export default App;
