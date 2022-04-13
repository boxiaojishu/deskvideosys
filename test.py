#!/usr/bin/python

import sqlite3

dn="192"
dip="192.168.0.129"
conn = sqlite3.connect('deskvideosys.db')
c = conn.cursor()
temp="UPDATE t_desks set deskip = '" + dip + "' where deskname='" + dn + "'"
print temp
c.execute(temp)
conn.commit()

cursor = conn.execute("SELECT id, deskname, deskloc, deskip  from t_desks")

for row in cursor:
   print "ID = ", row[0]
   print "NAME = ", row[1]
   print "ADDRESS = ", row[2]
   print "SALARY = ", row[3], "\n"

print "Operation done successfully"
conn.close()

