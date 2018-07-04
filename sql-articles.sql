USE [Articles]
GO

drop table Articles

CREATE TABLE [dbo].[Articles](
	[id] [int] primary key identity(1,1) NOT NULL,
	[idString] [nvarchar](512) NOT NULL,
	[title] [nvarchar](max) not null,
	[articleDate] [datetimeoffset](7) NOT NULL,
	[bodyText] [nvarchar](max) NOT NULL,
	[json] [nvarchar](max) NOT NULL,
	[lastUpdated] [datetimeoffset](7) NOT NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

alter table Articles add constraint uniqIdString unique(idString);

ALTER TABLE [dbo].[Articles] ADD  DEFAULT (getutcdate()) FOR [lastUpdated]
GO


CREATE TRIGGER trgArticlesLastUpdated
ON Articles
AFTER UPDATE AS
  UPDATE Articles
  SET [lastUpdated] = GETutcDATE()
  WHERE id IN (SELECT DISTINCT id FROM Inserted)
