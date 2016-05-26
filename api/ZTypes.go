/*
	Copyright (C) 2016  <Semchenko Aleksandr>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.If not, see <http://www.gnu.org/licenses/>.
*/

package api

type ZToken struct {
	Token string
}

type ZError struct {
	Error struct {
		      Code int
		      Text string
	      }
}

type ZTop struct {
	Page       int `json:"page"`
	PagesCount int `json:"pagesCount"`
	Tracks     []struct {
		Active                  bool    `json:"active"`
		ArtistID                int     `json:"artistId"`
		ArtistImageURLSquare100 string  `json:"artistImageUrlSquare100"`
		ArtistImageURLSquare250 string  `json:"artistImageUrlSquare250"`
		ArtistImageURLTop917    string  `json:"artistImageUrlTop917"`
		ArtistName              string  `json:"artistName"`
		Bitrate                 int     `json:"bitrate"`
		Block                   bool    `json:"block"`
		Count                   int     `json:"count"`
		Date                    int64   `json:"date"`
		Duration                string  `json:"duration"`
		HasRingBackTone         bool    `json:"hasRingBackTone"`
		ID                      int     `json:"id"`
		LastStamp               int     `json:"lastStamp"`
		Phantom                 bool    `json:"phantom"`
		Size                    float64 `json:"size"`
		Track                   string  `json:"track"`
		UserID                  int     `json:"userId"`
	} `json:"tracks"`
}

type ZSearch struct {
	Artist      struct {
			    About         string `json:"about"`
			    ID            int    `json:"id"`
			    ImageURI      string `json:"imageUri"`
			    Name          string `json:"name"`
			    SmallImageURI string `json:"smallImageUri"`
		    } `json:"artist"`
	Page        int      `json:"page"`
	PagesCount  int      `json:"pagesCount"`
	SuggestList []string `json:"suggestList"`
	Tracks      []struct {
		Active                  bool    `json:"active"`
		ArtistID                int     `json:"artistId"`
		ArtistImageURLSquare100 string  `json:"artistImageUrlSquare100"`
		ArtistImageURLSquare250 string  `json:"artistImageUrlSquare250"`
		ArtistImageURLTop917    string  `json:"artistImageUrlTop917"`
		ArtistName              string  `json:"artistName"`
		Bitrate                 int     `json:"bitrate"`
		Block                   bool    `json:"block"`
		Count                   int     `json:"count"`
		Date                    int     `json:"date"`
		Duration                string  `json:"duration"`
		HasRingBackTone         bool    `json:"hasRingBackTone"`
		ID                      int     `json:"id"`
		LastStamp               int     `json:"lastStamp"`
		Phantom                 bool    `json:"phantom"`
		Size                    float64 `json:"size"`
		Track                   string  `json:"track"`
		UserID                  int     `json:"userId"`
	} `json:"tracks"`
}

type ZGenre struct {
	Page       int `json:"page"`
	PagesCount int `json:"pagesCount"`
	Tracks     []struct {
		Active                  bool    `json:"active"`
		ArtistID                int     `json:"artistId"`
		ArtistImageURLSquare100 string  `json:"artistImageUrlSquare100"`
		ArtistImageURLSquare250 string  `json:"artistImageUrlSquare250"`
		ArtistImageURLTop917    string  `json:"artistImageUrlTop917"`
		ArtistName              string  `json:"artistName"`
		Bitrate                 int     `json:"bitrate"`
		Block                   bool    `json:"block"`
		Count                   int     `json:"count"`
		Date                    int     `json:"date"`
		Duration                string  `json:"duration"`
		HasRingBackTone         bool    `json:"hasRingBackTone"`
		ID                      int     `json:"id"`
		LastStamp               int     `json:"lastStamp"`
		Phantom                 bool    `json:"phantom"`
		Size                    float64 `json:"size"`
		Track                   string  `json:"track"`
		UserID                  int     `json:"userId"`
	} `json:"tracks"`
}

type ZArtist struct {
	Artist struct {
		       About         string `json:"about"`
		       ID            int    `json:"id"`
		       ImageURI      string `json:"imageUri"`
		       Name          string `json:"name"`
		       SmallImageURI string `json:"smallImageUri"`
	       } `json:"artist"`
}

type ZMusicSetList struct {
	List           []struct {
		About          string `json:"about"`
		CreateDate     int    `json:"createDate"`
		ID             int    `json:"id"`
		ImageURL       string `json:"imageUrl"`
		ImageURLTop917 string `json:"imageUrlTop917"`
		Name           string `json:"name"`
		PublishDate    int    `json:"publishDate"`
		SmallImageURL  string `json:"smallImageUrl"`
		TracksCount    int    `json:"tracksCount"`
		URL            string `json:"url"`
	} `json:"list"`
	MusicsetTypeID musicsetTypeID `json:"musicsetTypeId"`
}

type musicsetTypeID struct {
	Page       int `json:"page"`
	PagesCount int `json:"pagesCount"`
}

type ZMusicSetDetile struct {
	Musicset struct {
			 About          string `json:"about"`
			 CreateDate     int    `json:"createDate"`
			 ID             int    `json:"id"`
			 ImageURL       string `json:"imageUrl"`
			 ImageURLTop917 string `json:"imageUrlTop917"`
			 Name           string `json:"name"`
			 PublishDate    int    `json:"publishDate"`
			 SmallImageURL  string `json:"smallImageUrl"`
			 TracksCount    int    `json:"tracksCount"`
			 URL            string `json:"url"`
		 } `json:"musicset"`
	Tracks   []struct {
		ArtistID                int     `json:"artistId"`
		ArtistImageURLSquare100 string  `json:"artistImageUrlSquare100"`
		ArtistImageURLSquare250 string  `json:"artistImageUrlSquare250"`
		ArtistImageURLSquare800 string  `json:"artistImageUrlSquare800"`
		ArtistImageURLTop917    string  `json:"artistImageUrlTop917"`
		ArtistName              string  `json:"artistName"`
		Bitrate                 int     `json:"bitrate"`
		DlURL                   string  `json:"dlUrl"`
		Duration                string  `json:"duration"`
		FullName                string  `json:"fullName"`
		Ord                     int     `json:"ord"`
		PlayURL                 string  `json:"playUrl"`
		Size                    float64 `json:"size"`
		Track                   string  `json:"track"`
		TrackID                 int     `json:"trackId"`
	} `json:"tracks"`
}

type ZTrack struct {
	Rating float64 `json:"rating"`
	RbtURL string  `json:"rbtUrl"`
	Track  struct {
		       ArtistID                int      `json:"artistId"`
		       ArtistImageURLSquare100 string   `json:"artistImageUrlSquare100"`
		       ArtistImageURLSquare250 string   `json:"artistImageUrlSquare250"`
		       ArtistImageURLTop917    string   `json:"artistImageUrlTop917"`
		       ArtistName              string   `json:"artistName"`
		       Bitrate                 int      `json:"bitrate"`
		       Created                 int      `json:"created"`
		       Duration                int      `json:"duration"`
		       LyricAuthor             []string `json:"lyricAuthor"`
		       Lyrics                  struct {
						       Original []string `json:"original"`
					       } `json:"lyrics"`
		       MusicAuthor             []string `json:"musicAuthor"`
		       Name                    string   `json:"name"`
		       RightPossessors         []struct {
			       URL        string `json:"url"`
			       Name       string `json:"name"`
			       PictureURL string `json:"pictureUrl"`
		       } `json:"rightPossessors"`
		       Size                    float64 `json:"size"`
		       UserID                  int     `json:"userId"`
		       UserName                string  `json:"userName"`
	       } `json:"track"`
}

type ZTerms struct {
	Terms []string `json:"terms"`
}

type ZOptions struct {
	Options string `json:"options"`
}

type ZDownload struct {
	RbtURL string `json:"rbtUrl"`
	URL    string `json:"url"`
}

type ZPlay struct {
	RbtURL string `json:"rbtUrl"`
	URL    string `json:"url"`
}

type ZSettings struct {
	Token string `json:"token"`
}
