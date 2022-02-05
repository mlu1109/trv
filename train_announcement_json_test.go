package trv_test

const trainAnnouncementResponseJSON = `
{
	"RESPONSE": {
	  "RESULT": [
		{
		  "TrainAnnouncement": [
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc766c75398a",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:01:00.000+01:00",
			  "AdvertisedTrainIdent": "10139",
			  "Booking": [
				{
				  "Code": "BNA025",
				  "Description": "Hela tåget är obokat."
				}
			  ],
			  "Canceled": false,
			  "Deleted": false,
			  "Deviation": [
				{
				  "Code": "ANA055",
				  "Description": "Spårändrat"
				}
			  ],
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "Mälardalstrafik AB",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-02-05T15:01:17.867Z",
			  "NewEquipment": 0,
			  "Operator": "MTRN",
			  "OtherInformation": [
				{
				  "Code": "ONA180",
				  "Description": "Movingo gäller."
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA014",
				  "Description": "Mälartåg"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:02:00.000+01:00",
			  "TechnicalTrainIdent": "10139",
			  "TimeAtLocation": "2022-02-05T16:01:00.000+01:00",
			  "TimeAtLocationWithSeconds": "2022-02-05T16:01:06.000+01:00",
			  "ToLocation": [
				{
				  "LocationName": "Hpbg",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "13",
			  "TrainOwner": "MÄLAB",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Söö",
				  "Priority": 2,
				  "Order": 0
				},
				{
				  "LocationName": "Gn",
				  "Priority": 4,
				  "Order": 1
				},
				{
				  "LocationName": "Fle",
				  "Priority": 3,
				  "Order": 2
				},
				{
				  "LocationName": "K",
				  "Priority": 1,
				  "Order": 3
				}
			  ],
			  "WebLink": "https://malartag.se/",
			  "WebLinkName": "Mälartåg"
			},
			{
			  "ActivityId": "d6602b86-b939-ac5f-baad-45b6b5c436cc",
			  "ActivityType": "Avgang",
			  "Advertised": false,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:02:00.000+01:00",
			  "AdvertisedTrainIdent": "90575",
			  "Canceled": false,
			  "Deleted": false,
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-02-05T15:00:22.773Z",
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalTrainIdent": "90575",
			  "TimeAtLocation": "2022-02-05T15:59:00.000+01:00",
			  "TimeAtLocationWithSeconds": "2022-02-05T15:59:32.000+01:00",
			  "TrainOwner": "SJ"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc76963cc24f",
			  "ActivityType": "Avgang",
			  "Advertised": false,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:05:00.000+01:00",
			  "AdvertisedTrainIdent": "7846",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "A-train",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-02-05T15:06:02.015Z",
			  "NewEquipment": 0,
			  "Operator": "ATRAIN",
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:05:00.000+01:00",
			  "TechnicalTrainIdent": "7846",
			  "TimeAtLocation": "2022-02-05T16:05:00.000+01:00",
			  "TimeAtLocationWithSeconds": "2022-02-05T16:05:16.000+01:00",
			  "ToLocation": [
				{
				  "LocationName": "Arnn",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "1",
			  "TrainOwner": "ATRAIN",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Arns",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "WebLinkName": "A-train"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc7697beffc6",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:08:00.000+01:00",
			  "AdvertisedTrainIdent": "844",
			  "Canceled": true,
			  "Deleted": false,
			  "Deviation": [
				{
				  "Code": "ANA027",
				  "Description": "Inställt"
				},
				{
				  "Code": "ANA039",
				  "Description": "Nästa avgång"
				}
			  ],
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "MobileWebLink": "http://mobil.sj.se",
			  "ModifiedTime": "2022-01-29T02:43:02.750Z",
			  "NewEquipment": 0,
			  "Operator": "SJ",
			  "OtherInformation": [
				{
				  "Code": "ONA180",
				  "Description": "Movingo gäller."
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA025",
				  "Description": "SJ Regional"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:09:00.000+01:00",
			  "TechnicalTrainIdent": "844",
			  "ToLocation": [
				{
				  "LocationName": "U",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "x",
			  "TrainOwner": "SJ",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "WebLink": "http://www.sj.se/trafikinfo",
			  "WebLinkName": "SJ"
			},
			{
			  "ActivityId": "c933ff07-8ed8-e790-532b-fb828c2f4321",
			  "ActivityType": "Avgang",
			  "Advertised": false,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:08:00.000+01:00",
			  "AdvertisedTrainIdent": "90634",
			  "Canceled": false,
			  "Deleted": false,
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-02-05T15:05:14.597Z",
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalTrainIdent": "90634",
			  "TimeAtLocation": "2022-02-05T16:04:00.000+01:00",
			  "TimeAtLocationWithSeconds": "2022-02-05T16:04:59.000+01:00",
			  "TrainOwner": "SJ"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc769a632d81",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:11:00.000+01:00",
			  "AdvertisedTrainIdent": "950",
			  "Booking": [
				{
				  "Code": "BNA025",
				  "Description": "Hela tåget är obokat."
				}
			  ],
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "Mälardalstrafik AB",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-01-29T02:43:05.570Z",
			  "NewEquipment": 0,
			  "Operator": "MTRN",
			  "OtherInformation": [
				{
				  "Code": "ONA180",
				  "Description": "Movingo gäller."
				},
				{
				  "Code": "ONA151",
				  "Description": "Stannar ej vid Märsta."
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA014",
				  "Description": "Mälartåg"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:12:00.000+01:00",
			  "TechnicalTrainIdent": "20956",
			  "ToLocation": [
				{
				  "LocationName": "U",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "14",
			  "TrainComposition": [
				{
				  "Code": "TNA022",
				  "Description": "Långt tåg."
				}
			  ],
			  "TrainOwner": "MÄLAB",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Arnc",
				  "Priority": 1,
				  "Order": 0
				},
				{
				  "LocationName": "Kn",
				  "Priority": 2,
				  "Order": 1
				}
			  ],
			  "WebLink": "https://malartag.se/",
			  "WebLinkName": "Mälartåg"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc769367dd92",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:12:00.000+01:00",
			  "AdvertisedTrainIdent": "541",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "MobileWebLink": "http://mobil.sj.se",
			  "ModifiedTime": "2022-02-05T15:02:28.204Z",
			  "NewEquipment": 0,
			  "Operator": "SJ",
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA026",
				  "Description": "SJ Snabbtåg"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "Service": [
				{
				  "Code": "SNA001",
				  "Description": "Bistro"
				}
			  ],
			  "TechnicalDateTime": "2022-02-05T16:13:00.000+01:00",
			  "TechnicalTrainIdent": "541",
			  "ToLocation": [
				{
				  "LocationName": "Dk.kh",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "11",
			  "TrainComposition": [
				{
				  "Code": "TNA001",
				  "Description": "Vagnsordning 7-6-5-4-3-1"
				}
			  ],
			  "TrainOwner": "SJ",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Lp",
				  "Priority": 2,
				  "Order": 0
				},
				{
				  "LocationName": "Av",
				  "Priority": 3,
				  "Order": 1
				},
				{
				  "LocationName": "Hm",
				  "Priority": 4,
				  "Order": 2
				},
				{
				  "LocationName": "M",
				  "Priority": 1,
				  "Order": 3
				}
			  ],
			  "WebLink": "http://www.sj.se/trafikinfo",
			  "WebLinkName": "SJ"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc76952e0186",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:14:00.000+01:00",
			  "AdvertisedTrainIdent": "748",
			  "Booking": [
				{
				  "Code": "BNA001",
				  "Description": "Vagn 2 obokad."
				}
			  ],
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "MobileWebLink": "http://mobil.sj.se",
			  "ModifiedTime": "2022-02-05T14:42:14.703Z",
			  "NewEquipment": 0,
			  "Operator": "SJ",
			  "OtherInformation": [
				{
				  "Code": "ONA180",
				  "Description": "Movingo gäller."
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA025",
				  "Description": "SJ Regional"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:15:00.000+01:00",
			  "TechnicalTrainIdent": "748",
			  "ToLocation": [
				{
				  "LocationName": "Hpbg",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "15",
			  "TrainOwner": "SJ",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Bål",
				  "Priority": 3,
				  "Order": 0
				},
				{
				  "LocationName": "Ep",
				  "Priority": 4,
				  "Order": 1
				},
				{
				  "LocationName": "Vå",
				  "Priority": 1,
				  "Order": 2
				},
				{
				  "LocationName": "Ör",
				  "Priority": 2,
				  "Order": 3
				}
			  ],
			  "WebLink": "http://www.sj.se/trafikinfo",
			  "WebLinkName": "SJ"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc769253ec41",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:15:00.000+01:00",
			  "AdvertisedTrainIdent": "3943",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "Snälltåget",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-01-29T02:43:08.746Z",
			  "NewEquipment": 0,
			  "Operator": "HR",
			  "OtherInformation": [
				{
				  "Code": "ONA181",
				  "Description": "Movingo gäller ej."
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA069",
				  "Description": "Snälltåget"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "Service": [
				{
				  "Code": "SNA006",
				  "Description": "Krog"
				}
			  ],
			  "TechnicalDateTime": "2022-02-05T16:16:00.000+01:00",
			  "TechnicalTrainIdent": "3943",
			  "ToLocation": [
				{
				  "LocationName": "M",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "19",
			  "TrainComposition": [
				{
				  "Code": "TNA001",
				  "Description": "Vagnsordning 16-15-14-13-12-11-10"
				}
			  ],
			  "TrainOwner": "SNÄLL",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Nr",
				  "Priority": 1,
				  "Order": 0
				},
				{
				  "LocationName": "Lp",
				  "Priority": 2,
				  "Order": 1
				},
				{
				  "LocationName": "Hm",
				  "Priority": 3,
				  "Order": 2
				},
				{
				  "LocationName": "Lu",
				  "Priority": 4,
				  "Order": 3
				}
			  ],
			  "WebLink": "http://www.snalltaget.se",
			  "WebLinkName": "Snälltåget"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc769662e7e2",
			  "ActivityType": "Avgang",
			  "Advertised": false,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:20:00.000+01:00",
			  "AdvertisedTrainIdent": "7948",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "A-train",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-02-05T14:56:51.605Z",
			  "NewEquipment": 0,
			  "Operator": "ATRAIN",
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:20:00.000+01:00",
			  "TechnicalTrainIdent": "7948",
			  "ToLocation": [
				{
				  "LocationName": "Arnn",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "2",
			  "TrainOwner": "ATRAIN",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Arns",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "WebLinkName": "A-train"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc767f6644d9",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:21:00.000+01:00",
			  "AdvertisedTrainIdent": "20580",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "MobileWebLink": "http://mobil.sj.se",
			  "ModifiedTime": "2022-02-05T14:39:04.480Z",
			  "NewEquipment": 0,
			  "Operator": "SJ",
			  "OtherInformation": [
				{
				  "Code": "ONA099",
				  "Description": "Endast SJ Snabbtågsbiljetter gäller"
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA026",
				  "Description": "SJ Snabbtåg"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "Service": [
				{
				  "Code": "SNA001",
				  "Description": "Bistro"
				}
			  ],
			  "TechnicalDateTime": "2022-02-05T16:22:00.000+01:00",
			  "TechnicalTrainIdent": "20580",
			  "ToLocation": [
				{
				  "LocationName": "Uå",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "7",
			  "TrainComposition": [
				{
				  "Code": "TNA001",
				  "Description": "Vagnsordning 4-3-2-1"
				}
			  ],
			  "TrainOwner": "SJ",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Gä",
				  "Priority": 1,
				  "Order": 0
				},
				{
				  "LocationName": "Shv",
				  "Priority": 4,
				  "Order": 1
				},
				{
				  "LocationName": "Suc",
				  "Priority": 2,
				  "Order": 2
				},
				{
				  "LocationName": "Ök",
				  "Priority": 3,
				  "Order": 3
				}
			  ],
			  "WebLink": "http://www.sj.se/trafikinfo",
			  "WebLinkName": "SJ"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc76925fd822",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:29:00.000+01:00",
			  "AdvertisedTrainIdent": "411",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "MobileWebLink": "http://mobil.sj.se",
			  "ModifiedTime": "2022-01-29T02:43:18.994Z",
			  "NewEquipment": 0,
			  "Operator": "SJ",
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA026",
				  "Description": "SJ Snabbtåg"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "Service": [
				{
				  "Code": "SNA001",
				  "Description": "Bistro"
				}
			  ],
			  "TechnicalDateTime": "2022-02-05T16:30:00.000+01:00",
			  "TechnicalTrainIdent": "411",
			  "ToLocation": [
				{
				  "LocationName": "G",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "13",
			  "TrainComposition": [
				{
				  "Code": "TNA001",
				  "Description": "Vagnsordning 1-3-4-5-6-7"
				}
			  ],
			  "TrainOwner": "SJ",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "K",
				  "Priority": 2,
				  "Order": 0
				},
				{
				  "LocationName": "Sk",
				  "Priority": 1,
				  "Order": 1
				},
				{
				  "LocationName": "Hr",
				  "Priority": 3,
				  "Order": 2
				}
			  ],
			  "WebLink": "http://www.sj.se/trafikinfo",
			  "WebLinkName": "SJ"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc76963cc259",
			  "ActivityType": "Avgang",
			  "Advertised": false,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:35:00.000+01:00",
			  "AdvertisedTrainIdent": "7848",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "A-train",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-01-29T02:43:23.920Z",
			  "NewEquipment": 0,
			  "Operator": "ATRAIN",
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:35:00.000+01:00",
			  "TechnicalTrainIdent": "7848",
			  "ToLocation": [
				{
				  "LocationName": "Arnn",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "1",
			  "TrainOwner": "ATRAIN",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Arns",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "WebLinkName": "A-train"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc766e38b7ef",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:37:00.000+01:00",
			  "AdvertisedTrainIdent": "10846",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "MobileWebLink": "http://mobil.sj.se",
			  "ModifiedTime": "2022-01-29T02:43:27.183Z",
			  "NewEquipment": 0,
			  "Operator": "SJ",
			  "OtherInformation": [
				{
				  "Code": "ONA180",
				  "Description": "Movingo gäller."
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA025",
				  "Description": "SJ Regional"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:38:00.000+01:00",
			  "TechnicalTrainIdent": "10846",
			  "ToLocation": [
				{
				  "LocationName": "U",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "3",
			  "TrainOwner": "SJ",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Mr",
				  "Priority": 1,
				  "Order": 0
				},
				{
				  "LocationName": "Kn",
				  "Priority": 2,
				  "Order": 1
				}
			  ],
			  "WebLink": "http://www.sj.se/trafikinfo",
			  "WebLinkName": "SJ"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc76943e8c1d",
			  "ActivityType": "Avgang",
			  "Advertised": false,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:37:00.000+01:00",
			  "AdvertisedTrainIdent": "7061",
			  "Booking": [
				{
				  "Code": "BNA002",
				  "Description": "Vagn 12 och 14 obokade."
				}
			  ],
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "Tågab",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-01-29T02:43:27.292Z",
			  "NewEquipment": 0,
			  "Operator": "TÅGAB",
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA034",
				  "Description": "Tågab"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "Service": [
				{
				  "Code": "SNA005",
				  "Description": "Servering"
				}
			  ],
			  "TechnicalDateTime": "2022-02-05T16:37:00.000+01:00",
			  "TechnicalTrainIdent": "7061",
			  "ToLocation": [
				{
				  "LocationName": "Ksc",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "17a",
			  "TrainOwner": "TÅGAB",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "K",
				  "Priority": 1,
				  "Order": 0
				},
				{
				  "LocationName": "Hpbg",
				  "Priority": 2,
				  "Order": 1
				},
				{
				  "LocationName": "Dg",
				  "Priority": 4,
				  "Order": 2
				},
				{
				  "LocationName": "Khn",
				  "Priority": 3,
				  "Order": 3
				}
			  ],
			  "WebLink": "http://www.tagakeriet.se",
			  "WebLinkName": "Tågab"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc7685f535a8",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:39:00.000+01:00",
			  "AdvertisedTrainIdent": "277",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Gä",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "MobileWebLink": "http://mobil.sj.se",
			  "ModifiedTime": "2022-01-29T02:42:06.946Z",
			  "NewEquipment": 0,
			  "Operator": "SJ",
			  "OtherInformation": [
				{
				  "Code": "ONA181",
				  "Description": "Movingo gäller ej."
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA023",
				  "Description": "SJ InterCity"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:40:00.000+01:00",
			  "TechnicalTrainIdent": "277",
			  "ToLocation": [
				{
				  "LocationName": "Lp",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "12b",
			  "TrainOwner": "SJ",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Flb",
				  "Priority": 2,
				  "Order": 0
				},
				{
				  "LocationName": "Söö",
				  "Priority": 3,
				  "Order": 1
				},
				{
				  "LocationName": "Nr",
				  "Priority": 1,
				  "Order": 2
				}
			  ],
			  "WebLink": "http://www.sj.se/trafikinfo",
			  "WebLinkName": "SJ"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc767f6b0997",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:45:00.000+01:00",
			  "AdvertisedTrainIdent": "2061",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "MTRN",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-01-29T02:43:32.326Z",
			  "NewEquipment": 0,
			  "Operator": "MTRN",
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA068",
				  "Description": "MTRX"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:46:00.000+01:00",
			  "TechnicalTrainIdent": "2061",
			  "ToLocation": [
				{
				  "LocationName": "G",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "14a",
			  "TrainComposition": [
				{
				  "Code": "TNA001",
				  "Description": "Vagnsordning A-B-C-D-E"
				},
				{
				  "Code": "TNA032",
				  "Description": "Café i vagn B"
				}
			  ],
			  "TrainOwner": "MTRN",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "K",
				  "Priority": 2,
				  "Order": 0
				},
				{
				  "LocationName": "Sk",
				  "Priority": 1,
				  "Order": 1
				},
				{
				  "LocationName": "A",
				  "Priority": 3,
				  "Order": 2
				}
			  ],
			  "WebLink": "https://mtrx.travel/sv",
			  "WebLinkName": "MTRX"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc7696654a26",
			  "ActivityType": "Avgang",
			  "Advertised": false,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:50:00.000+01:00",
			  "AdvertisedTrainIdent": "7950",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "A-train",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-01-29T02:43:36.151Z",
			  "NewEquipment": 0,
			  "Operator": "ATRAIN",
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:50:00.000+01:00",
			  "TechnicalTrainIdent": "7950",
			  "ToLocation": [
				{
				  "LocationName": "Arnn",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "2",
			  "TrainOwner": "ATRAIN",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Arns",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "WebLinkName": "A-train"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc767f79578c",
			  "ActivityType": "Avgang",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:53:00.000+01:00",
			  "AdvertisedTrainIdent": "20955",
			  "Booking": [
				{
				  "Code": "BNA025",
				  "Description": "Hela tåget är obokat."
				}
			  ],
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "U",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "Mälardalstrafik AB",
			  "LocationSignature": "Cst",
			  "ModifiedTime": "2022-01-29T02:43:05.255Z",
			  "NewEquipment": 0,
			  "Operator": "MTRN",
			  "OtherInformation": [
				{
				  "Code": "ONA180",
				  "Description": "Movingo gäller."
				},
				{
				  "Code": "ONA120",
				  "Description": "Bakre delen: slutstation Stockholm C."
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA014",
				  "Description": "Mälartåg"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "TechnicalDateTime": "2022-02-05T16:54:00.000+01:00",
			  "TechnicalTrainIdent": "20955",
			  "ToLocation": [
				{
				  "LocationName": "Arb",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "12",
			  "TrainOwner": "MÄLAB",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Söö",
				  "Priority": 3,
				  "Order": 0
				},
				{
				  "LocationName": "Lg",
				  "Priority": 4,
				  "Order": 1
				},
				{
				  "LocationName": "Sgs",
				  "Priority": 2,
				  "Order": 2
				},
				{
				  "LocationName": "Et",
				  "Priority": 1,
				  "Order": 3
				}
			  ],
			  "WebLink": "https://malartag.se/",
			  "WebLinkName": "Mälartåg"
			},
			{
			  "ActivityId": "1500adde-0a5d-4065-08d9-dc7693f1ee05",
			  "ActivityType": "Ankomst",
			  "Advertised": true,
			  "AdvertisedTimeAtLocation": "2022-02-05T16:59:00.000+01:00",
			  "AdvertisedTrainIdent": "598",
			  "Canceled": false,
			  "Deleted": false,
			  "EstimatedTimeIsPreliminary": false,
			  "FromLocation": [
				{
				  "LocationName": "Cst",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "InformationOwner": "SJ",
			  "LocationSignature": "Cst",
			  "MobileWebLink": "http://mobil.sj.se",
			  "ModifiedTime": "2022-01-29T02:43:41.076Z",
			  "NewEquipment": 0,
			  "Operator": "SJ",
			  "OtherInformation": [
				{
				  "Code": "ONA099",
				  "Description": "Endast SJ Snabbtågsbiljetter gäller"
				}
			  ],
			  "PlannedEstimatedTimeAtLocationIsValid": false,
			  "ProductInformation": [
				{
				  "Code": "PNA026",
				  "Description": "SJ Snabbtåg"
				}
			  ],
			  "ScheduledDepartureDateTime": "2022-02-05T00:00:00.000+01:00",
			  "Service": [
				{
				  "Code": "SNA001",
				  "Description": "Bistro"
				}
			  ],
			  "TechnicalDateTime": "2022-02-05T17:00:00.000+01:00",
			  "TechnicalTrainIdent": "598",
			  "ToLocation": [
				{
				  "LocationName": "Ös",
				  "Priority": 1,
				  "Order": 0
				}
			  ],
			  "TrackAtLocation": "4",
			  "TrainComposition": [
				{
				  "Code": "TNA001",
				  "Description": "Vagnsordning 4-3-2-1"
				}
			  ],
			  "TrainOwner": "SJ",
			  "TypeOfTraffic": [
				{
				  "Code": "YNA001",
				  "Description": "Tåg"
				}
			  ],
			  "ViaToLocation": [
				{
				  "LocationName": "Gä",
				  "Priority": 1,
				  "Order": 0
				},
				{
				  "LocationName": "Bn",
				  "Priority": 2,
				  "Order": 1
				},
				{
				  "LocationName": "Ls",
				  "Priority": 3,
				  "Order": 2
				},
				{
				  "LocationName": "Åg",
				  "Priority": 4,
				  "Order": 3
				}
			  ],
			  "WebLink": "http://www.sj.se/trafikinfo",
			  "WebLinkName": "SJ"
			}
		  ]
		}
	  ]
	}
  }
`
