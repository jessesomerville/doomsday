from rich.console import Console
from rich.table import Table

t = Table()

t.add_column('Month')
t.add_column('Day')
t.add_column('Mnemonic')
t.add_column('All Days')

t.add_row('Jan', '1/3, 1/4', '3rd in 3, 4th in 4', '3, 10, 17, 24, 31 (+1 for leap)')
t.add_row('Feb', '2/28, 2/29', 'last day of Feb', '0, 7, 14, 21, 28 (+1 for leap)')
t.add_row('Mar', '3/14', 'last day of Feb', '0, 7, 14, 21, 28')
t.add_row('Apr', '4/4', 'evens double', '4, 11, 18, 25, 32')
t.add_row('May', '5/9', '9-to-5 at 7-11', '2, 9, 16, 23, 30')
t.add_row('Jun', '6/6', 'evens double', '6, 13, 20, 27')
t.add_row('Jul', '7/11', '9-to-5 at 7-11', '4, 11, 18, 25, 32')
t.add_row('Aug', '8/8', 'evens double', '1, 8, 15, 22, 29')
t.add_row('Sep', '9/5', '9-to-5 at 7-11', '5, 12, 19, 26')
t.add_row('Oct', '10/10', 'evens double', '3, 10, 17, 24, 31')
t.add_row('Nov', '11/7', '9-to-5 at 7-11', '0, 7, 14, 21, 28')
t.add_row('Dec', '12/12', 'evens double', '5, 12, 19, 26')

console = Console()
console.print(t)
