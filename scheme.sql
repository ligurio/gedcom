DROP TABLE IF EXISTS `famchild`;
CREATE TABLE `famchild` (
`famID` varchar(40) NOT NULL DEFAULT '',
`child` varchar(40) NOT NULL DEFAULT '',
PRIMARY KEY (`famID`,`child`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `family`;
CREATE TABLE `family` (
`famID` varchar(40) NOT NULL DEFAULT '',
`husband` varchar(40) DEFAULT NULL,
`wife` varchar(40) DEFAULT NULL,
`marr_date` varchar(255) DEFAULT NULL,
`marr_plac` varchar(255) DEFAULT NULL,
`marr_sour` varchar(255) DEFAULT NULL,
`marb_date` varchar(255) DEFAULT NULL,
`marb_plac` varchar(255) DEFAULT NULL,
`marb_sour` varchar(255) DEFAULT NULL,
PRIMARY KEY (`famID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `person_st`;
CREATE TABLE `person_st` (
`persID` varchar(40) NOT NULL DEFAULT '',
`name` varchar(255) DEFAULT NULL,
`vorname` varchar(255) DEFAULT NULL,
`marname` varchar(255) DEFAULT NULL,
`sex` char(1) DEFAULT NULL,
`birt_date` varchar(255) DEFAULT NULL,
`birt_plac` varchar(255) DEFAULT NULL,
`birt_sour` varchar(255) DEFAULT NULL,
`taufe_date` varchar(255) DEFAULT NULL,
`taufe_plac` varchar(255) DEFAULT NULL,
`taufe_sour` varchar(255) DEFAULT NULL,
`deat_date` varchar(255) DEFAULT NULL,
`deat_plac` varchar(255) DEFAULT NULL,
`deat_sour` varchar(255) DEFAULT NULL,
`buri_date` varchar(255) DEFAULT NULL,
`buri_plac` varchar(255) DEFAULT NULL,
`buri_sour` varchar(255) DEFAULT NULL,
`occupation` varchar(255) DEFAULT NULL,
`occu_date` varchar(255) DEFAULT NULL,
`occu_plac` varchar(255) DEFAULT NULL,
`occu_sour` varchar(255) DEFAULT NULL,
`religion` varchar(80) DEFAULT NULL,
`confi_date` varchar(255) DEFAULT NULL,
`confi_plac` varchar(255) DEFAULT NULL,
`confi_sour` varchar(255) DEFAULT NULL,
`note` longtext,
PRIMARY KEY (`persID`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
