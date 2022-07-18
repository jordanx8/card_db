import React from 'react';
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
  Button,
  TableToolbarContent,
  TableSelectAll,
  TableToolbarAction
} from '@carbon/react';
import './App.scss';

function App() {
  const rows2 = [
    {
      id: 'a',
      firstname: "Steven",
      lastname: "Adams",
      seasonsplayed: "( 2020-21 )"
    }
  ]
  const headers2 = [
    {
      key: 'firstname',
      header: "First Name"
    },
    {
      key: 'lastname',
      header: "Last Name"
    },
    {
      key: 'seasonsplayed',
      header: "Season"
    },
  ]
  const rows3 = [
    {
      id: 'a',
      season: "2019-20 ( Oklahoma City Thunder )",
      manufacturer: "Panini",
      set: "NBA Hoops Premium Stock",
      insert: "",
      parallel: "Box Set Pulsar Prizm",
      number: "130",
      notes: null
    },
    {
      id: 'b',
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
      key: 'season',
      header: 'Season',
    },
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
                  {row.cells.map((cell) => (
                    <TableCell key={cell.id}>{cell.value}</TableCell>
                  ))}
                </TableExpandRow>
                {row.isExpanded && (
                  <TableExpandedRow colSpan={headers.length + 1}>
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
        </>
      )}
    </DataTable>
  );
}

export default App;
