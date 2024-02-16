package provider

const vastTemplate = `<VAST version="3.0" xmlns:xs="http://www.w3.org/2001/XMLSchema">
    <Ad id="{{ad_id}}">
        <InLine>
            <AdSystem version="4.0">iabtechlab</AdSystem>
            <AdTitle>iabtechlab video ad</AdTitle>
            <Error>http://example.com/error</Error>
            <Impression id="Impression-ID">{{base_url}}/vast/impression/{{ad_id}}</Impression>
            <Creatives>
                <Creative id="{{ad_id}}" sequence="1">
                    <Linear>
                        <Duration>00:00:16</Duration>

                        <TrackingEvents>
                            <Tracking event="start">{{base_url}}/vast/start/{{ad_id}}</Tracking>
                            <Tracking event="firstQuartile">{{base_url}}/vast/firstQuartile/{{ad_id}}</Tracking>
                            <Tracking event="midpoint">{{base_url}}/vast/midpoint/{{ad_id}}</Tracking>
                            <Tracking event="thirdQuartile">{{base_url}}/vast/thirdQuartile/{{ad_id}}</Tracking>
                            <Tracking event="complete">{{base_url}}/vast/complete/{{ad_id}}</Tracking>
                        </TrackingEvents>

                        <VideoClicks>
                            <ClickThrough id="blog">
                                <![CDATA[{{destination_url}}]]>
                            </ClickThrough>
                           <ClickTracking>
                                <![CDATA[{{base_url}}/vast/click/{{ad_id}}]]>
                            </ClickTracking>
                        </VideoClicks>

                        <MediaFiles>
                            <MediaFile id="{{ad_id}}" delivery="progressive" type="video/mp4" bitrate="500" width="400" height="300" minBitrate="360" maxBitrate="1080" scalable="1" maintainAspectRatio="1" codec="0">
                                <![CDATA[{{video_url}}]]>
                            </MediaFile>
                        </MediaFiles>

                    </Linear>
                </Creative>
            </Creatives>
        </InLine>
    </Ad>
</VAST>`
