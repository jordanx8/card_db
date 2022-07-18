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
import './App.scss';

function App() {
  const images =
  {
    Steven_Adams: "https://www.basketball-reference.com/req/202106291/images/players/adamsst01.jpg"
  }

  const headers1 = [
    {
      key: 'season',
      header: 'Season'
    },
  ]

  const rows1 = [
    {
      id: '1',
      season: "2019-20 ( Oklahoma City Thunder )",
    },
    {
      id: '2',
      season: "2020-21 ( New Orleans Pelicans )",
    }
  ]

  const rows2 = [
    {
      id: '1',
      firstname: "Steven",
      lastname: "Adams",
      seasonsplayed: "( 2020-21 )",
      numcards: "4"
    }
  ]
  const headers2 = [
    {
      key: 'firstname',
      header: "First"
    },
    {
      key: 'lastname',
      header: "Last"
    },
    {
      key: 'seasonsplayed',
      header: "Season"
    },
    {
      key: 'numcards',
      header: "Number of Cards"
    }
  ]
  const rows3 = [
    {
      id: '1',
      season: "2019-20 ( Oklahoma City Thunder )",
      manufacturer: "Panini",
      set: "NBA Hoops Premium Stock",
      insert: "",
      parallel: "Box Set Pulsar Prizm",
      number: "130",
      notes: null
    },
    {
      id: '2',
      season: "2020-21 ( New Orleans Pelicans )",
      manufacturer: "Panini",
      set: "Prizm",
      insert: "",
      parallel: "Box Set Pulsar Prizm",
      number: "110",
      notes: null
    },
  ];
  const headers3 = [
    {
      key: 'manufacturer',
      header: 'Manufacturer',
    },
    {
      key: 'set',
      header: 'Set',
    },
    {
      key: 'insert',
      header: 'Insert',
    },
    {
      key: 'parallel',
      header: 'Parallel',
    },
    {
      key: 'number',
      header: 'Card Number',
    },
    {
      key: 'notes',
      header: 'Notes',
    },
  ];

  return (
    <DataTable rows={rows2} headers={headers2}>
      {({ rows, headers, getTableProps, getHeaderProps, getRowProps, getBatchActionProps, onInputChange }) => (
        <>
          <TableToolbar>
            <TableToolbarContent>
              <TableToolbarSearch
                tabIndex={getBatchActionProps().shouldShowBatchActions ? -1 : 0}
                onChange={onInputChange}
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
                    <TableCell><img src={images[row.cells[0].value + '_' + row.cells[1].value]} alt={row.cells[0].value + ' ' + row.cells[1].value}></img></TableCell>
                    {row.cells.map((cell) => (
                      <TableCell key={cell.id}>{cell.value}</TableCell>
                    ))}
                  </TableExpandRow>
                  {row.isExpanded && (
                    <TableExpandedRow colSpan={headers.length + 2}>
                      <DataTable rows={rows1} headers={headers1}>
                        {({ rows, headers, getTableProps, getHeaderProps, getRowProps }) => (
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
                                      <TableCell key={cell.id}>{cell.value}</TableCell>
                                    ))}
                                  </TableExpandRow>
                                  {row.isExpanded && (
                                    <TableExpandedRow colSpan={headers.length + 2}>
                                      <DataTable rows={rows3} headers={headers3}>
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
