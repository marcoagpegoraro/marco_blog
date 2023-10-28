package services_test

import (
	"testing"

	"github.com/marcoagpegoraro/marco_blog/services"
)

func TestGetImageFromString(t *testing.T) {
	postBody := `
    <p>A expressão Lorem ipsum em design gráfico e editoração é um texto padrão em latim utilizado na produção gráfica para
    preencher os espaços de texto em publicações para testar e ajustar aspectos visuais antes de utilizar conteúdo real.
    <a href="https://pt.wikipedia.org/wiki/Lorem_ipsum" target="_blank">Wikipédia</a></p>
    <p><br></p>
    <p><img src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEASABIAAD/4gIoSUNDX1BST0ZJTEUAAQEAAAIYAAAAAAQwAABtbnRyUkdCIFhZWiAAAAAAAAAAAAAAAABhY3NwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAA9tYAAQAAAADTLQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlkZXNjAAAA8AAAAHRyWFlaAAABZAAAABRnWFlaAAABeAAAABRiWFlaAAABjAAAABRyVFJDAAABoAAAAChnVFJDAAABoAAAAChiVFJDAAABoAAAACh3dHB0AAAByAAAABRjcHJ0AAAB3AAAADxtbHVjAAAAAAAAAAEAAAAMZW5VUwAAAFgAAAAcAHMAUgBHAEIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAiVglG2jPViodZyMvJWydjNwL7WwM45lGXWKrxLqjS8dq/7Emym2VVgKkziuYmcK9Qtf34INGtxOS3LfXzzGUULgSVj5F6mGBX2g6aMb8zKZGBdtcHRDkckyhqV2Kmcua7QF7IDGur2qMgoBd7Xl94bJ1OU6zLbW/xEHZU33qLEwUBiC8Ch2iiTkXMtrMmZ17hunpggli5vnEDVOV/MEoQ4Xpy0eRCRiGBMaJiuuRQte7t4hoCg1XV7QfyIlu2HlGMi392obVf72AWTEUFP42OAC/PssCi6e+lKK7JEDsthYiMiuX5XqWBGJMyMI39RO9QwjUVHQYYndjAcM0i4yshYIvYYzQEjAFDAuUeygYXJAIUDU7jI4O0NCgIhCTBwQMSWAoMCzqEpQt+pUE0lAoti82AisloFzwOQ4oW4uTYKBAWgNcguJOshIi13sosoUGOQLWcgCwSLxzwU3sLo0YsgCHmBiCh1Bk0EBddqgjJYTGBtIttwbJA5LC+BtQz7QFqpzrYWlSFoKShcgNQcoj0AUnBnwW4dpNhpl+ZAm6aJAiiTaveQHfEaDDSN0gDodH6jqvRFVSLFeki1R9c21UYwxxr6/DWIBjR0+OxAD7uLoGAW0tYCyUFqwuqiCe3GH9WPQIHpV1H2AEJmX6ZIXEbMCY2VoobeSP2CJGD53gW65Jg4cM2eDNrGg/slygIIa6mj7CC9q6W17J+VS8+8GQ6eogQ77JJPtdjOccdWDidGERAGwwapShgMAqVKC9CRl1R6kjD//2Q==">
    </p>
    <p><br></p>
    <p>A expressão Lorem ipsum em design gráfico e editoração é um texto padrão em latim utilizado na produção gráfica para
      preencher os espaços de texto em publicações para testar e ajustar aspectos visuais antes de utilizar conteúdo real.
      <a href="https://pt.wikipedia.org/wiki/Lorem_ipsum" target="_blank">Wikipédia</a></p>
    <p><br></p>
    <p><img src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wgARCAcUBQADASIAAhEBAxEB/8QAHAAAAgMBAQEBAAAAAAAAAAAAAAECAwQFBgcI/8QAGgEBAQEBAQEBAAAAAAAAAAAAAAECAwQFBv/aAAwDAQACEAMQAAAB8omfpPioYIZCGUmAw8/4z6n5fWfCldn6v8uAayAAAABx1Neb3oGJSQhoTGKSDuR4px36no+GfPfoedzzpj13D5s47/Aa3ipN9ClEJoaRBiAEBKSi7NGnBoznUVWY1Wro2UFqShWrSpyiDgWCBYqSWIxRiJOuVk3U4miSJtxIcpUWGdRk2RFBHEjRbVOW7Tj053svp0cesoTpzZOE4UNEVor0OzNK+qxKNOpZXVTrOmuo1JykS1RuSUFsbKSxVQWwuYwkEWRqx1hZAcQchVY3LBSiklAVsmsVZOKb5W49ImvZ5AD0fOAAaBxkpe59G+IfVPzX6Pi+14vD83o3dry/tkAFPGex+SdOdEg/U/mQDWQAAAAOSM83uSaoTIRJEWxUxpEZKOMrAAJRC2I2aiSaUTBOH0SoTubw/RgYg2ZiYj6HoQ16MJ0gxFykFMdHoDe5g0RxVTibPaOv7+WJVtizBIGaiirmFpvBLyzqWXbCoGhgSqgGoFfki1rzNoydShspbERqXhZ85RFszNHtMEN3LVliGOr2zPFWwXQqDtXgwEtDIqFlqHEENdwCS9z//2Q==">
    </p>
    <p><br></p>
    <p>dsadsadsa</p>
    <p><img src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/4gHYSUNDX1BST0ZJTEUAAQEAAAHIbGNtcwIQAABtbnRyUkdCIFhZWiAH4gADABQACQAOAB1hY3NwTVNGVAAAAABzYXdzY3RybAAAAAAAAAAAAAAAAAAA9tYAAQAAAADTLWhhbmSdkQA9QICwPUB0LIGepSKOAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlkZXNjAAAA8AAAAF9jcHJ0AAABDAAAAAx3dHB0AAABGAAAABRyWFlaAAABLAAAABRnWFlaAAABQAAAABRiWFlaAAABVAAAABRyVFJDAAABaAAAAGBnVFJDAAABaAAAAGBiVFJDAAABaAAAAGBkZXNjAAAAAAAAAAV1UkdCAAAAAAAAAAAAAAAAdGV4dAAAAABDQzAAWFlaIAAAAAAAAPNUAAEAAAABFslYWVogAAAAAAAAb6AAADjyAAADj1hZWiAAAAAAAABilgAAt4kAABjaWFlaIAAAAAAAACSgAAAPhQAAtsRjdXJ2AAAAAAAAACoAAAB8APgBnZO2DQAkjweK/CWpX0llbWt/pm1/MtY/KWVCcdKsx3dro3w90e/TTLae8eaVFknQMBz3Hfiq0iW/hXwpqFi97bXmo6kAm23bzFiQE9Tiq+pXcEnw50W2SeMzxTyF41YFuvUj0oAt6jcW2u/D2TV2sLO2u7W6WINbReWrKe2Kv2Fill4NS/wBDsbW9mLr9pkvE3bfkH3QfesOyvLdPhVf2nnxpcPexssPmAsRxztzVzSrO1ksopvDWurp92mBPDcz+WH4HTjHWgDO8S6no2oWtsIdO+w6shIufKUCN+g4rnB713fjbUNOm0GytpLi2vdZQ/v54Oh9efyrhdpUUAJRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAf//Z">
    </p>
    <p>Test</p>
  `

	images := services.AWSS3Service.GetBase64ImagesFromString(postBody)

	if images == nil || len(images) != 3 {
		t.Errorf("Length of images correct")
	}

	for _, image := range images {
		if image[0:4] != "data" {
			t.Errorf("base64 image not valid")
		}
	}
}
