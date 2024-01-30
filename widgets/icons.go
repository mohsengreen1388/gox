package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const Size = 94

type Icons rl.Vector2

var NoIcon Icons = Icons{0, 0}
var ArrowLeft = Icons{1, 0}
var ArrowRight = Icons{Size * 1, 0}
var ArrowDown = Icons{Size * 2, 0}
var ArrowUp = Icons{Size * 3, 0}
var Abc = Icons{Size * 4, 0}
var AccountBalanceWallet = Icons{Size * 5, 0}
var AccountBalance = Icons{Size * 6, 0}
var AccountCircle = Icons{Size * 7, 0}
var Add = Icons{Size * 8, 0}
var Apps = Icons{Size * 9, 0}
var ArrowBack = Icons{Size * 10, 0}
var ArrowForward = Icons{1, Size}
var Article = Icons{Size * 1, Size}
var Attachment = Icons{Size * 2, Size}
var Backup = Icons{Size * 3, Size}
var Bookmark = Icons{Size * 4, Size}
var Brush = Icons{Size * 5, Size}
var CheckCircle = Icons{Size * 6, Size}
var Checklist = Icons{Size * 7, Size}
var Check = Icons{Size * 8, Size}
// var nop = Icons{Size * 9, Size}
var Trash = Icons{Size * 10, Size}
var Close = Icons{1, Size * 2}
var Cloud = Icons{Size * 1, Size * 2}
var Copy = Icons{Size * 2, Size * 2}
var Past = Icons{Size * 3, Size * 2}
var CopyRight = Icons{Size * 4, Size * 2}
var CreditCard = Icons{Size * 5, Size * 2}
var Description = Icons{Size * 6, Size * 2}
var DownloadDone = Icons{Size * 7, Size * 2}
var Download = Icons{Size * 8, Size * 2}
var EditNote = Icons{Size * 9, Size * 2}
var Edit = Icons{Size * 10, Size * 2}
var DarkMode = Icons{1, Size * 3}
var Dashboard = Icons{Size * 1, Size * 3}
var Eject = Icons{Size * 2, Size * 3}
var EmojiEmotions = Icons{Size * 3, Size * 3}
var EventAvailable = Icons{Size * 4, Size * 3}
var ExitToApp = Icons{Size * 5, Size * 3}
var FileUpload = Icons{Size * 6, Size * 3}
var Fingerprint = Icons{Size * 7, Size * 3}
var List = Icons{Size * 8, Size * 3}
var FullScreen = Icons{Size * 9, Size * 3}
var Star = Icons{Size * 10, Size * 3}
var ExpandLessExpandMore = Icons{1, Size * 4}
var ExpandLess = Icons{Size * 1, Size * 4}
var FavoriteBorder = Icons{Size * 2, Size * 4}
var Favorite = Icons{Size * 3, Size * 4}
var FileDownload = Icons{Size * 4, Size * 4}
var FilePresent = Icons{Size * 5, Size * 4}
var FirstPage = Icons{Size * 6, Size * 4}
var FolderCopy = Icons{Size * 7, Size * 4}
var Folder = Icons{Size * 8, Size * 4}
var Game = Icons{Size * 9, Size * 4}
var Grid = Icons{Size * 10, Size * 4}
var Translate = Icons{1, Size * 5}
var HeadPhones = Icons{Size * 1, Size * 5}
var Help = Icons{Size * 2, Size * 5}
var History = Icons{Size * 3, Size * 5}
var ImageFile = Icons{Size * 4, Size * 5}
var Info = Icons{Size * 5, Size * 5}
var Inputs = Icons{Size * 6, Size * 5}
var KeyboardHide = Icons{Size * 7, Size * 5}
var Keyboard = Icons{Size * 8, Size * 5}
var Language = Icons{Size * 9, Size * 5}
var LastPage = Icons{Size * 10, Size * 5}

var Lightbulb = Icons{1, Size * 6}
//var nop = Icons{Size * 1, Size * 6}
var Location = Icons{Size * 2, Size * 6}
var Login = Icons{Size * 3, Size * 6}
var ManageAccounts = Icons{Size * 4, Size * 6}
var Menu = Icons{Size * 5, Size * 6}
var OpenInBrowser = Icons{Size * 6, Size * 6}
var AddPost = Icons{Size * 7, Size * 6}
var Preview = Icons{Size * 8, Size * 6}
var Redeem = Icons{Size * 9, Size * 6}
var RemoveDone = Icons{Size * 10, Size * 6}
var Maximize = Icons{1, Size * 7}
var MoreVert = Icons{Size * 1, Size * 7}
var Movie = Icons{Size * 2, Size * 7}
var Music = Icons{Size * 3, Size * 7}
var AddNote = Icons{Size * 4, Size * 7}
var Paid = Icons{Size * 5, Size * 7}
var PermMedia = Icons{Size * 6, Size * 7}
var Printer = Icons{Size * 7, Size * 7}
var Public = Icons{Size * 8, Size * 7}
var RemoveRedEye = Icons{Size * 9, Size * 7}
var Save = Icons{Size * 10, Size * 7}
var NavigateBefore = Icons{1, Size * 8}
var NavigateNext = Icons{Size * 1, Size * 8}
var NotInterested = Icons{Size * 2, Size * 8}
var Pause = Icons{Size * 3, Size * 8}
var Percent = Icons{Size * 4, Size * 8}
var PersonAdd = Icons{Size * 5, Size * 8}
var Camrea = Icons{Size * 6, Size * 8}
var Reply = Icons{Size * 7, Size * 8}
var RestorOfTrash = Icons{Size * 8, Size * 8}
var RoketLunch = Icons{Size * 9, Size * 8}
var Savings = Icons{Size * 10, Size * 8}
var Play = Icons{1, Size * 9}
var Search = Icons{Size * 1, Size * 9}
var Send = Icons{Size * 2, Size * 9}
var Text = Icons{Size * 3, Size * 9}
var Message = Icons{Size * 4, Size * 9}
var UnLike = Icons{Size * 5, Size * 9}
var Like = Icons{Size * 6, Size * 9}
var Video = Icons{Size * 7, Size * 9}
var AddVideo = Icons{Size * 8, Size * 9}
var Share = Icons{Size * 9, Size * 9}
var Market = Icons{Size * 10, Size * 9}
var Source = Icons{1, Size * 10}
var Sort = Icons{Size * 1, Size * 10}
var TableView = Icons{Size * 2, Size * 10}
var ViewList = Icons{Size * 3, Size * 10}
var Zoom = Icons{Size * 4, Size * 10}
var SkipNext = Icons{Size * 5, Size * 10}
var SkipPrevious = Icons{Size * 6, Size * 10}
var VolumeOff = Icons{Size * 7, Size * 10}
var VolumeUp = Icons{Size * 8, Size * 10}
var Home = Icons{Size * 9, Size * 10}
var Setting = Icons{Size * 10, Size * 10}

type Icon struct {
	dataFileIcon   []byte
	iconDataDecode []byte
	image          *rl.Image
	texture        rl.Texture2D
}

func (ic *Icon) Draw(icon Icons, x, y float32, iconSize int8, color rl.Color) {
	ic.body(icon, x, y, iconSize, color)
}

func (ic *Icon) body(icon Icons, x, y float32, iconSize int8, color rl.Color) {
	iconSize = ic.checkSize(iconSize)
	rl.DrawTexturePro(ic.texture, rl.Rectangle{icon.X, icon.Y, Size, Size}, rl.Rectangle{x, y, float32(iconSize), float32(iconSize)},
		rl.Vector2{0, 0},
		0, color)
}

func (ic *Icon) checkSize(size int8) int8 {
	if size > 94 {
		return 94
	}
	return size
}

func (ic *Icon) Unload() {
	rl.UnloadTexture(ic.texture)
}
