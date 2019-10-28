CREATE TABLE ConfigTable (RecID INTEGER PRIMARY KEY, RecType INTEGER, Title TEXT, DataRec BLOB );
CREATE INDEX idxRecType ON ConfigTable (RecType);
CREATE TABLE PersonTable (PersonID INTEGER PRIMARY KEY, UniqueID TEXT, Sex INTEGER, EditDate FLOAT, ParentID INTEGER, SpouseID INTEGER, Color INTEGER, Relate1 INTEGER, Relate2 INTEGER, Flags INTEGER, Living INTEGER, IsPrivate INTEGER, Proof INTEGER, Bookmark INTEGER, Note BLOB );
CREATE TABLE FamilyTable (FamilyID INTEGER PRIMARY KEY, FatherID INTEGER, MotherID INTEGER, ChildID INTEGER, HusbOrder INTEGER, WifeOrder INTEGER, IsPrivate INTEGER, Proof INTEGER, SpouseLabel INTEGER, FatherLabel INTEGER, MotherLabel INTEGER, Note BLOB );
CREATE INDEX idxFamilyFatherID ON FamilyTable (FatherID);
CREATE INDEX idxFamilyMotherID ON FamilyTable (MotherID);
CREATE TABLE ChildTable (RecID INTEGER PRIMARY KEY, ChildID INTEGER, FamilyID INTEGER, RelFather INTEGER, RelMother INTEGER, ChildOrder INTEGER, IsPrivate INTEGER, ProofFather INTEGER, ProofMother INTEGER, Note BLOB );
CREATE INDEX idxChildID ON ChildTable (ChildID);
CREATE INDEX idxChildFamilyID ON ChildTable (FamilyID);
CREATE INDEX idxChildOrder ON ChildTable (ChildOrder);
CREATE TABLE EventTable (EventID INTEGER PRIMARY KEY, EventType INTEGER, OwnerType INTEGER, OwnerID INTEGER, FamilyID INTEGER, PlaceID INTEGER, SiteID INTEGER, Date TEXT, SortDate INTEGER, IsPrimary INTEGER, IsPrivate INTEGER, Proof INTEGER, Status INTEGER, EditDate FLOAT, Sentence BLOB, Details BLOB, Note BLOB );
CREATE INDEX idxOwnerEvent ON EventTable (OwnerID,EventType);
CREATE INDEX idxOwnerDate ON EventTable (OwnerID,SortDate);
CREATE TABLE AddressTable (AddressID INTEGER PRIMARY KEY, AddressType INTEGER, Name TEXT COLLATE RMNOCASE, Street1 TEXT, Street2 TEXT, City TEXT, State TEXT, Zip TEXT, Country TEXT, Phone1 TEXT, Phone2 TEXT, Fax TEXT, Email TEXT, URL TEXT, Latitude INTEGER, Longitude INTEGER, Note BLOB );
CREATE INDEX idxAddressName ON AddressTable (Name);
CREATE TABLE FactTypeTable (FactTypeID INTEGER PRIMARY KEY, OwnerType INTEGER, Name TEXT COLLATE RMNOCASE, Abbrev TEXT, GedcomTag TEXT, UseValue INTEGER, UseDate INTEGER, UsePlace INTEGER, Sentence BLOB, Flags INTEGER );
CREATE INDEX idxFactTypeName ON FactTypeTable (Name);
CREATE INDEX idxFactTypeAbbrev ON FactTypeTable (Abbrev);
CREATE INDEX idxFactTypeGedcomTag ON FactTypeTable (GedcomTag);
CREATE TABLE MultimediaTable (MediaID INTEGER PRIMARY KEY, MediaType INTEGER, MediaPath TEXT, MediaFile TEXT COLLATE RMNOCASE, URL TEXT, Thumbnail BLOB, Caption TEXT COLLATE RMNOCASE, RefNumber TEXT COLLATE RMNOCASE, Date TEXT, SortDate INTEGER, Description BLOB );
CREATE INDEX idxMediaFile ON MultimediaTable (MediaFile);
CREATE INDEX idxMediaURL ON MultimediaTable (URL);
CREATE TABLE MediaLinkTable (LinkID INTEGER PRIMARY KEY, MediaID INTEGER, OwnerType INTEGER, OwnerID INTEGER, IsPrimary INTEGER, Include1 INTEGER, Include2 INTEGER, Include3 INTEGER, Include4 INTEGER, SortOrder INTEGER, RectLeft INTEGER, RectTop INTEGER, RectRight INTEGER, RectBottom INTEGER, Note TEXT, Caption TEXT COLLATE RMNOCASE, RefNumber TEXT COLLATE RMNOCASE, Date TEXT, SortDate INTEGER, Description BLOB );
CREATE INDEX idxMediaOwnerID ON MediaLinkTable (OwnerID);
CREATE INDEX idxMediaCaption ON MediaLinkTable (Caption);
CREATE TABLE NameTable (NameID INTEGER PRIMARY KEY, OwnerID INTEGER, Surname TEXT COLLATE RMNOCASE, Given TEXT COLLATE RMNOCASE, Prefix TEXT COLLATE RMNOCASE, Suffix TEXT COLLATE RMNOCASE, Nickname TEXT COLLATE RMNOCASE, NameType INTEGER, Date TEXT, SortDate INTEGER, IsPrimary INTEGER, IsPrivate INTEGER, Proof INTEGER, EditDate FLOAT, Sentence BLOB, Note BLOB, BirthYear INTEGER, DeathYear INTEGER );
CREATE INDEX idxNameOwnerID ON NameTable (OwnerID);
CREATE INDEX idxSurname ON NameTable (Surname);
CREATE INDEX idxGiven ON NameTable (Given);
CREATE INDEX idxSurnameGiven ON NameTable (Surname, Given, BirthYear, DeathYear);
CREATE INDEX idxNamePrimary ON NameTable (IsPrimary);
CREATE TABLE PlaceTable (PlaceID INTEGER PRIMARY KEY, PlaceType INTEGER, Name TEXT COLLATE RMNOCASE, Abbrev TEXT, Normalized TEXT, Latitude INTEGER, Longitude INTEGER, LatLongExact INTEGER, MasterID INTEGER, Note BLOB );
CREATE INDEX idxPlaceName ON PlaceTable (Name);
CREATE INDEX idxPlaceAbbrev ON PlaceTable (Abbrev);
CREATE TABLE ResearchTable (TaskID INTEGER PRIMARY KEY, TaskType INTEGER, OwnerID INTEGER, OwnerType INTEGER, RefNumber TEXT, Name TEXT COLLATE RMNOCASE, Status INTEGER, Priority INTEGER, Date1 TEXT, Date2 TEXT, Date3 TEXT, SortDate1 INTEGER, SortDate2 INTEGER, SortDate3 INTEGER, Filename TEXT, Details BLOB );
CREATE INDEX idxResearchOwnerID ON ResearchTable (OwnerID);
CREATE INDEX idxResearchName ON ResearchTable (Name);
CREATE TABLE ResearchItemTable (ItemID INTEGER PRIMARY KEY, LogID INTEGER, Date TEXT, SortDate INTEGER, RefNumber TEXT, Repository TEXT, Goal TEXT, Source TEXT, Result TEXT );
CREATE INDEX idxResearchItemLogID ON ResearchItemTable (LogID);
CREATE TABLE SourceTable (SourceID INTEGER PRIMARY KEY, Name TEXT COLLATE RMNOCASE, RefNumber TEXT, ActualText TEXT, Comments TEXT, IsPrivate INTEGER, TemplateID INTEGER, Fields BLOB );
CREATE INDEX idxSourceName ON SourceTable (Name);
CREATE TABLE CitationTable (CitationID INTEGER PRIMARY KEY, OwnerType INTEGER, SourceID INTEGER, OwnerID INTEGER, Quality TEXT, IsPrivate INTEGER, Comments BLOB, ActualText BLOB, RefNumber TEXT, Flags INTEGER, Fields BLOB );
CREATE INDEX idxCitationSourceID ON CitationTable (SourceID);
CREATE INDEX idxCitationOwnerID ON CitationTable (OwnerID);
CREATE TABLE AddressLinkTable (LinkID INTEGER PRIMARY KEY, OwnerType INTEGER, AddressID INTEGER, OwnerID INTEGER, AddressNum INTEGER, Details TEXT );
CREATE TABLE WitnessTable (WitnessID INTEGER PRIMARY KEY, EventID INTEGER, PersonID INTEGER, WitnessOrder INTEGER, Role INTEGER, Sentence TEXT, Note BLOB, Given TEXT COLLATE RMNOCASE, Surname TEXT COLLATE RMNOCASE, Prefix TEXT COLLATE RMNOCASE, Suffix TEXT COLLATE RMNOCASE );
CREATE INDEX idxWitnessEventID ON WitnessTable (EventID);
CREATE INDEX idxWitnessPersonID ON WitnessTable (PersonID);
CREATE TABLE LinkTable (LinkID INTEGER PRIMARY KEY, extSystem INTEGER, LinkType INTEGER, rmID INTEGER, extID TEXT, Modified INTEGER, extVersion TEXT, extDate FLOAT, Status INTEGER, Note BLOB );
CREATE INDEX idxLinkRmId ON LinkTable (rmID);
CREATE INDEX idxLinkExtId ON LinkTable (extID);
CREATE TABLE LinkAncestryTable (LinkID INTEGER PRIMARY KEY, extSystem INTEGER, LinkType INTEGER, rmID INTEGER, extID TEXT, Modified INTEGER, extVersion TEXT, extDate FLOAT, Status INTEGER, Note BLOB );
CREATE INDEX idxLinkAncestryRmId ON LinkAncestryTable (rmID);
CREATE INDEX idxLinkAncestryExtId ON LinkAncestryTable (extID);
CREATE TABLE RoleTable (RoleID INTEGER PRIMARY KEY, RoleName TEXT COLLATE RMNOCASE, EventType INTEGER, RoleType INTEGER, Sentence TEXT );
CREATE INDEX idxRoleEventType ON RoleTable (EventType);
CREATE TABLE GroupTable (RecID INTEGER PRIMARY KEY, GroupID INTEGER, StartID INTEGER, EndID INTEGER );
CREATE TABLE ExclusionTable (RecID INTEGER PRIMARY KEY, ExclusionType INTEGER, ID1 INTEGER, ID2 INTEGER );
CREATE UNIQUE INDEX idxExclusionIndex ON ExclusionTable (ExclusionType, ID1, ID2);
CREATE TABLE SourceTemplateTable (TemplateID INTEGER PRIMARY KEY, Name TEXT COLLATE RMNOCASE, Description TEXT, Favorite INTEGER, Category TEXT, Footnote TEXT, ShortFootnote TEXT, Bibliography TEXT, FieldDefs BLOB );
CREATE INDEX idxSourceTemplateName ON SourceTemplateTable (Name);
CREATE TABLE LabelTable (LabelID INTEGER PRIMARY KEY, LabelType INTEGER, LabelValue INTEGER, LabelName TEXT COLLATE RMNOCASE, Description TEXT );
CREATE INDEX idxLabelType ON LabelTable (LabelType);
CREATE TABLE URLTable (LinkID INTEGER PRIMARY KEY, OwnerType INTEGER, OwnerID INTEGER, LinkType INTEGER, Name TEXT, URL TEXT, Note BLOB );
CREATE INDEX idxURLOwnerID ON URLTable (OwnerID);
