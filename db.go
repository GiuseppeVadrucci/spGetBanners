package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)





--
-- Table structure for table `banners`
--

CREATE TABLE `banners` (
  `zoneId` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `widht`  int(11) NOT NULL,
  `height` int(11) NOT NULL,
  `banners`, JSON_OBJECT(
	    "id": "",
        "date": "",
        "background_color": "",
        "background_image": "",
        "button_text": "",
        "button_color": "",
        "button_background": "",
        "title": "",
        "description": "",
        "text_align": "",
        "link": "https:"",
        "link_isExternal"": 
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `banners`
--

INSERT INTO `banners` (`zoneId`,`widht`,`height`,`banners`) VALUES

--
-- Procedures
--

DELIMITER $$
CREATE PROCEDURE [cms].spGetBanners(@LanguageCode VARCHAR(30),@PageId INT,@DeviceTypeId TINYINT,@ErrorMessage VARCHAR(1000))
BEGIN
       SELECT * FROM banners where zoneId = @LanguageCode and id = @PageId;
END $$
DELIMITER ;

