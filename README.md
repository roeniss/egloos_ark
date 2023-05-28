# 이글루스의 방주 (egloos ark) 프로젝트

2023 년 6 월 중순, 이글루스가 완전 폐쇄를 공지를 올렸습니다. 동시에 "폐쇄 후 연말까지 백업 플랜을 제시하겠다"고 얘기했습니다.

이글루스는 일찍이 그늘에 묻힌, 그들만의 평화로운 공간이 된 지 오래입니다. 웰컴투동막골과 다를 바가 없지요. 또한 많은 블로그 주인장들이 게시물이 10000개를 넘기고 있습니다. 도저히 손으로 옮길 수 있는 수준이 아닙니다.

**그러므로, 혹시라도 이글루스 운영진의 플랜이 제대로 이행되지 않으면, 이글루스의 여러 유저들이 자신의 게시물을 고스란히 잃게 됩니다.** 그때가서 대처하면 늦습니다. 마치 싸이월드 때처럼요.

이 프로젝트는 일찍이 제가 좋아하는 한 이글루스 주인장을 위해 개발되었으나 널리 퍼졌으면 좋겠다는 마음으로 이렇게 공유하게 되었습니다. 이글루스에 연이 없어서 어떻게 널리 퍼뜨려야 할 지 모르겠습니다만, 가능하다면 알고 계신 이글루스 유저분들께 공유해주시면 감사하겠습니다.

## 이 글을 보신 분이 블로그 주인장이신데 사용법을 모르시겠다면

저에게 메일로 블로그 주소 보내주세요 (roeniss2@gmail.com). 제가 일단 긁어서 저장하고 있겠습니다. 이 데이터를 쓸지 말지는 나중에 이글루스 측의 대처를 보고 판단하시지요.

## 이 프로젝트로 할 수 있는 것

- 특정 블로그의 게시물 몽땅 스크랩해오기
  - 내용을 이쁘게 담는 것이 아니라, 본문에 해당하는 html 을 통째로 뜯어냅니다. 이거로 뭘 할 건지는 아래 "로드맵" 섹션을 참고해주세요.
  - 가져오는 것: 본문 (이미지들의 url 포함), 글 작성 날짜 (초 단위가 아니라 분 단위로 찍힐 수도 있음), 태그들, 댓글들, 제목, 카테고리
  - 이글루스는 블로그 테마 (템플릿) 가 꽤 자유로운 편입니다. 저는 서두에 말씀드렸듯 아주 특정한 누군가의 블로그만 바라보고 만들었기 때문에 생각처럼 잘 안 긁어질 가능성이 있습니다.
  - 비공개 글은 가져오지 못합니다.
  - 가져온 데이터는 sqlite db 파일로 만들어지기 때문에 비개발자가 "잘 가져왔나" 확인하는 것은 어려울 것입니다.

## 사용법

이 레포를 다운받고 루트 디렉토리에서 다음 명령어를 실행합니다.

```bash 

go run github.com/roeniss/egloos_ark/cli crawl -b $BLOG_ID

```

## 로드맵

- 이미지 모두 다운로드 (이글루스가 닫히면 더이상 이미지 서빙도 안 할 것 같음) - 현재 작업중
- 특정 플랫폼에 전부 업로드 - 이 부분은 이글루스 팀의 대처를 보고 할지말지 결정

## 유사한 프로그램

윈도우에서 실행할 수 있고, 본인 계정으로 접속하기 때문에 비공개 글까지 수집 가능한 프로그램이 있습니다. http://orumi.egloos.com/7580955 한 번 확인해주시길 바랍니다.

