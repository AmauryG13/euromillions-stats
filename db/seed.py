#!/usr/bin/env python

import os
import sqlite3
import csv
from io import TextIOWrapper
from zipfile import ZipFile 

# Setup path
cwd = os.path.dirname(__file__)

dbDIR = cwd
dbFILENAME = "euromillions.sqlite"
dbFILE = os.path.join(cwd, '..', dbDIR, dbFILENAME)

dataDIR = os.path.join(cwd, '../data')
dataARCHIVES = [
    'euromillions_201402.zip',
    'euromillions_201105.zip',
    'euromillions_200402.zip',
]

# Setup database
sqltableHISTORY = """
CREATE TABLE IF NOT EXISTS history 
(
    id INT PRIMARY KEY,
    date DATE,
    ball1 INT,
    ball2 INT,
    ball3 INT,
    ball4 INT,
    ball5 INT,
    star1 INT,
    star2 INT
)
"""

sqlinsertHISTORY = """
INSERT INTO history (date, ball1, ball2, ball3, ball4, ball4, star1, star2)
VALUES {values}
"""

def initDatabase():
    conn = sqlite3.connect(dbFILE)
    cursor = conn.cursor()
    cursor.execute(sqltableHISTORY)

    return cursor

# Deals with data
class Draw:
    def __init__(self, date, balls, stars):
        self.date = date
        self.balls = balls
        self.stars = stars

    def insertStrings(self):
        return """
            (
                {date}, {balls[0]}, {balls[1]}, {balls[2]}, {balls[3]}, {balls[4]}, {stars[0]}, {stars[1]}
            )
        """.format(date=self.date, balls=self.balls, stars=self.stars)

def readData(archiveIdx = 0):
    data = []
    archiveFILENAME = dataARCHIVES[archiveIdx]
    archiveFILE = os.path.join(dataDIR, archiveFILENAME)

    with ZipFile(archiveFILE) as zfile:
        filename = zfile.namelist()

        with zfile.open(filename[0]) as file:
            reader = csv.reader(TextIOWrapper(file, 'utf-8'))

            for row in reader:
                items = (row[0]).split(';')
                print(items)
                data.append(
                    Draw(
                        date = items[2],
                        balls = [items[4], items[5], items[6], items[7], items[8]],
                        stars = [items[9], items[10]]
                    ))

    return data

def insertDraw2DB(cursor, data):
    for draw in data:
        cursor.execute(
            sqlinsertHISTORY.replace("{values}", draw.insertStrings())
        )

if __name__ == "__main__":
    cursor = initDatabase()
    
    data = readData()
    insertDraw2DB(cursor, data)