<?php namespace Edged\Mulungu\Manager;

use Symfony\Component\Finder\Finder;
use Symfony\Component\Finder\SplFileInfo;
use Symfony\Component\Filesystem\Filesystem;

class FileManager{
	
	public static $FILE = "File";
	public static $DIRECTORY = "Directory";
	public static $CONTENT_FORMAT_JSON = "json";
	public static $CONTENT_FORMAT_RAW = "raw";
	private $tree;
	private $fileSystem;
	private $cache;
	private $cachePath;
	
		
	public function __construct(){
		$this->fileSystem = new Filesystem();
	}
	
	public function getDirectoryContents( $dir, $tree=[], $parent="#", $depth=0, $ignores=["*.php","robots.txt"] ){
		
		$finder = new Finder();
		$finder->in( $dir );
		$finder->depth( $depth );
		
		foreach ( $ignores as $ignore ){ $finder->notName( $ignore ); }
		
		foreach ( $finder as $item ) {
			
			$itemId = md5(  $item->getBasename() );
			$node = [
					
				"id"=>$itemId,
				"parent"=>$parent,
				"name" => $item->getBasename(),
				"realPath" => $item->getRealpath(),
				"type"=>$this->getType( $item ),
				"extension"=>$item->getExtension(),
				"icon"=>$this->getIconByExtension( $item->getExtension() ),
				"editor"=>$this->getEditorByExtension( $item->getExtension() )
			];
			
			$children = ( $this->getType( $item ) == self::$DIRECTORY )?$this->getDirectoryContents( $item->getRealpath(), [], $itemId ):null;
			$node[ "children" ] = $children;
			
			$tree[] = $node;
		}
		
		$this->tree = $tree;
		return $this->tree;
	}
	
	public function getType( SplFileInfo $splFileInfo ){
		if( $splFileInfo->isDir() ){
			return self::$DIRECTORY;
		}else if( $splFileInfo->isFile() ){
			return self::$FILE;
		}
		
		return "Unknown";
	}
	
	public function getById( $id ){
		if( empty( $this->tree ) )
			return null;
	
		return $this->search( "id", $id, $this->tree );
	}
	
	public function search( $key, $value, $target ){
		foreach ( $target as $candidate ){
			if( $candidate[ $key ] === $value ){
				return $candidate;
			}else{
				$candidate = $this->search($key, $value, $candidate[ "children" ] );
				if( $candidate ){
					return $candidate;
				}
			}
		}
	
	}	
	
	public function getIconByExtension( $extension ){
		switch( $extension ){
			case "css": return "fa-css3"; break;
			case "txt": return "fa-file-text"; break;
			case "png": return "fa-file-image-o"; break;
			case "jpg": return "fa-file-image-o"; break;
			case "php": return "fa-code"; break;
			case "js": return "fa-code"; break;
			case "eot": return "fa-font"; break;
			case "ttf": return "fa-font"; break;
			case "otf": return "fa-font"; break;
			case "woff": return "fa-font"; break;
			default: return "fa-file-o";
		}
	}
	
	public function getEditorByExtension( $extension ){
		switch( $extension ){
			case "css": return "code"; break;
			case "txt": return "code"; break;
			case "png": return "image"; break;
			case "jpg": return "image"; break;
			case "php": return "code"; break;
			case "js": return "code"; break;
			case "eot": return "font"; break;
			case "ttf": return "font"; break;
			case "otf": return "font"; break;
			case "woff": return "font"; break;
			default: return "unknown";
		}
	}
	
	public function guessLocationByExtension( $extension ){
		switch( $extension ){
			case "png":
			case "jpg":
				return "img";
				break;
			case "eot":
			case "ttf":
			case "otf":
			case "woff":
				return "fonts";
				break;
			case "css":
				return "css";
				break;
			case "js":
				return "js";
				break;
			default: return null; break;
		}
	}
	
	public function getExtension( $fileName ){
		return substr( $fileName, strrpos($fileName, ".")+1, strlen($fileName) );
	}	
	
	public function content( $filePath, $format="raw" ){
		$content = null;
		
		if( file_exists( $filePath ) )
			$content = file_get_contents( $filePath );
		
		switch( $format ){
			case self::$CONTENT_FORMAT_JSON : $content = json_decode( $content, true ); break;
			default: break;
		} 
		
		return $content;
	}
	
	public function dump( $content, $path, $permission=0777 ){
		$this->fileSystem->dumpFile( $path, $content, $permission );
	}
}