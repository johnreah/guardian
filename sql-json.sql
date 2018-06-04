
create table test(col1 int, col2 nvarchar(max))

alter table test add constraint jsonConstraint check (isjson(col2)=1)

declare @json nvarchar(max)

set @json = 
N'
{
	"id" : 2,
	"info": { 
		"name": "John", 
		"surname": "Smith" 
	}, 
	"age": 25 
}
'

insert into test values(1, @json)

select * from test

select col1, json_value(col2, '$.info.surname') from test

drop table test
