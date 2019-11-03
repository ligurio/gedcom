        DROP TABLE IF EXISTS note;
CREATE TABLE note (
                  handle CHARACTER(25) PRIMARY KEY,
                  gid    CHARACTER(25),
                  text   TEXT,
                  format INTEGER,
                  note_type1   INTEGER,
                  note_type2   TEXT,
                  change INTEGER,
                  private BOOLEAN);
        DROP TABLE IF EXISTS name;
CREATE TABLE name (
                  handle CHARACTER(25) PRIMARY KEY,
                  primary_name BOOLEAN,
                  private BOOLEAN, 
                  first_name TEXT, 
                  suffix TEXT, 
                  title TEXT, 
                  name_type0 INTEGER, 
                  name_type1 TEXT, 
                  group_as TEXT, 
                  sort_as INTEGER,
                  display_as INTEGER, 
                  call TEXT,
                  nick TEXT,
                  famnick TEXT);
        DROP TABLE IF EXISTS surname;
CREATE TABLE surname (
                  handle CHARACTER(25),
                  surname TEXT, 
                  prefix TEXT, 
                  primary_surname BOOLEAN, 
                  origin_type0 INTEGER,
                  origin_type1 TEXT,
                  connector TEXT);
CREATE INDEX idx_surname_handle ON 
                  surname(handle);
        DROP TABLE IF EXISTS date;
CREATE TABLE date (
                  handle CHARACTER(25) PRIMARY KEY,
                  calendar INTEGER, 
                  modifier INTEGER, 
                  quality INTEGER,
                  day1 INTEGER, 
                  month1 INTEGER, 
                  year1 INTEGER, 
                  slash1 BOOLEAN,
                  day2 INTEGER, 
                  month2 INTEGER, 
                  year2 INTEGER, 
                  slash2 BOOLEAN,
                  text TEXT, 
                  sortval INTEGER, 
                  newyear INTEGER);
        DROP TABLE IF EXISTS person;
CREATE TABLE person (
                  handle CHARACTER(25) PRIMARY KEY,
                  gid CHARACTER(25), 
                  gender INTEGER, 
                  death_ref_handle TEXT, 
                  birth_ref_handle TEXT, 
                  change INTEGER, 
                  private BOOLEAN);
        DROP TABLE IF EXISTS family;
CREATE TABLE family (
                 handle CHARACTER(25) PRIMARY KEY,
                 gid CHARACTER(25), 
                 father_handle CHARACTER(25), 
                 mother_handle CHARACTER(25), 
                 the_type0 INTEGER, 
                 the_type1 TEXT, 
                 change INTEGER, 
                 private BOOLEAN);
        DROP TABLE IF EXISTS place;
CREATE TABLE place (
                 handle CHARACTER(25) PRIMARY KEY,
                 gid CHARACTER(25), 
                 title TEXT, 
                 value TEXT, 
                 the_type0 INTEGER,
                 the_type1 TEXT,
                 code TEXT,
                 long TEXT, 
                 lat TEXT, 
                 lang TEXT, 
                 change INTEGER, 
                 private BOOLEAN);
        DROP TABLE IF EXISTS place_ref;
CREATE TABLE place_ref (
                   handle             CHARACTER(25) PRIMARY KEY,
                   from_place_handle  CHARACTER(25),
                   to_place_handle    CHARACTER(25));
        DROP TABLE IF EXISTS place_name;
CREATE TABLE place_name (
                  handle        CHARACTER(25) PRIMARY KEY,
                  from_handle   CHARACTER(25),
                  value         CHARACTER(25),
                  lang          CHARACTER(25));
        DROP TABLE IF EXISTS event;
CREATE TABLE event (
                 handle CHARACTER(25) PRIMARY KEY,
                 gid CHARACTER(25), 
                 the_type0 INTEGER, 
                 the_type1 TEXT, 
                 description TEXT, 
                 change INTEGER, 
                 private BOOLEAN);
        DROP TABLE IF EXISTS citation;
CREATE TABLE citation (
                 handle CHARACTER(25) PRIMARY KEY,
                 gid CHARACTER(25), 
                 confidence INTEGER,
                 page CHARACTER(25),
                 source_handle CHARACTER(25),
                 change INTEGER,
                 private BOOLEAN);
        DROP TABLE IF EXISTS source;
CREATE TABLE source (
                 handle CHARACTER(25) PRIMARY KEY,
                 gid CHARACTER(25), 
                 title TEXT, 
                 author TEXT, 
                 pubinfo TEXT, 
                 abbrev TEXT, 
                 change INTEGER,
                 private BOOLEAN);
        DROP TABLE IF EXISTS media;
CREATE TABLE media (
                 handle CHARACTER(25) PRIMARY KEY,
                 gid CHARACTER(25), 
                 path TEXT, 
                 mime TEXT, 
                 desc TEXT,
                 checksum INTEGER,
                 change INTEGER, 
                 private BOOLEAN);
        DROP TABLE IF EXISTS repository_ref;
CREATE TABLE repository_ref (
                 handle CHARACTER(25) PRIMARY KEY,
                 ref CHARACTER(25), 
                 call_number TEXT, 
                 source_media_type0 INTEGER,
                 source_media_type1 TEXT,
                 private BOOLEAN);
        DROP TABLE IF EXISTS repository;
CREATE TABLE repository (
                 handle CHARACTER(25) PRIMARY KEY,
                 gid CHARACTER(25), 
                 the_type0 INTEGER, 
                 the_type1 TEXT,
                 name TEXT, 
                 change INTEGER, 
                 private BOOLEAN);
        DROP TABLE IF EXISTS link;
CREATE TABLE link (
                 from_type CHARACTER(25), 
                 from_handle CHARACTER(25), 
                 to_type CHARACTER(25), 
                 to_handle CHARACTER(25));
CREATE INDEX idx_link_to ON 
                  link(from_type, from_handle, to_type);
        DROP TABLE IF EXISTS markup;
CREATE TABLE markup (
                 handle CHARACTER(25) PRIMARY KEY,
                 markup0 INTEGER, 
                 markup1 TEXT, 
                 value INTEGER, 
                 start_stop_list TEXT);
        DROP TABLE IF EXISTS event_ref;
CREATE TABLE event_ref (
                 handle CHARACTER(25) PRIMARY KEY,
                 ref CHARACTER(25), 
                 role0 INTEGER, 
                 role1 TEXT, 
                 private BOOLEAN);
        DROP TABLE IF EXISTS person_ref;
CREATE TABLE person_ref (
                 handle CHARACTER(25) PRIMARY KEY,
                 description TEXT,
                 private BOOLEAN);
        DROP TABLE IF EXISTS child_ref;
CREATE TABLE child_ref (
                 handle CHARACTER(25) PRIMARY KEY,
                 ref CHARACTER(25), 
                 frel0 INTEGER,
                 frel1 CHARACTER(25),
                 mrel0 INTEGER,
                 mrel1 CHARACTER(25),
                 private BOOLEAN);
        DROP TABLE IF EXISTS lds;
CREATE TABLE lds (
                 handle CHARACTER(25) PRIMARY KEY,
                 type INTEGER, 
                 place CHARACTER(25), 
                 famc CHARACTER(25), 
                 temple TEXT, 
                 status INTEGER, 
                 private BOOLEAN);
        DROP TABLE IF EXISTS media_ref;
CREATE TABLE media_ref (
                 handle CHARACTER(25) PRIMARY KEY,
                 ref CHARACTER(25),
                 role0 INTEGER,
                 role1 INTEGER,
                 role2 INTEGER,
                 role3 INTEGER,
                 private BOOLEAN);
        DROP TABLE IF EXISTS address;
CREATE TABLE address (
                handle CHARACTER(25) PRIMARY KEY,
                private BOOLEAN);
        DROP TABLE IF EXISTS location;
CREATE TABLE location (
                 handle CHARACTER(25) PRIMARY KEY,
                 street TEXT, 
                 locality TEXT,
                 city TEXT, 
                 county TEXT, 
                 state TEXT, 
                 country TEXT, 
                 postal TEXT, 
                 phone TEXT,
                 parish TEXT);
        DROP TABLE IF EXISTS attribute;
CREATE TABLE attribute (
                 handle CHARACTER(25) PRIMARY KEY,
                 the_type0 INTEGER, 
                 the_type1 TEXT, 
                 value TEXT, 
                 private BOOLEAN);
        DROP TABLE IF EXISTS url;
CREATE TABLE url (
                 handle CHARACTER(25) PRIMARY KEY,
                 path TEXT, 
                 desc TXT, 
                 type0 INTEGER,
                 type1 TEXT,                  
                 private BOOLEAN);
        DROP TABLE IF EXISTS datamap;
CREATE TABLE datamap (
                 from_handle CHARACTER(25),
                 the_type0 INTEGER,
                 the_type1 TEXT,
                 value_field TXT,
                 private BOOLEAN);
        DROP TABLE IF EXISTS tag;
CREATE TABLE tag (
                 handle CHARACTER(25) PRIMARY KEY,
                 name TEXT,
                 color TEXT,
                 priority INTEGER,
                 change INTEGER);
