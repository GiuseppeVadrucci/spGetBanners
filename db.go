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
	    "id": 1,
        "date": "2019-09-13T08:00:00Z",
        "background_color": "#333333",
        "background_image": "https://cdn.betstarters.com/banners/banner-1920x450.jpg",
        "button_text": "Play Now!",
        "button_color": "#FFFFFF",
        "button_background": "#CDCDCD",
        "title": "Play Now!",
        "description": "Play Now!",
        "text_align": "left",
        "link": "https:\/\/online.sbbet.me\/banner-sport\/baner-kosarka-polufinale-1920x450\/",
        "link_isExternal": true
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

