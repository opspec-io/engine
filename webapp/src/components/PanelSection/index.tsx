import React from 'react'
import brandColors from '../../brandColors'
import { css, cx } from '@emotion/css'

const commonCss = css({
  backgroundColor: brandColors.reallyReallyLightGray,
  paddingRight: '1.2rem',
  paddingLeft: '1.2rem',
  width: '100%',
  verticalAlign: 'middle'
})

/**
 * A titled section of the panel containing rows & optionally a summary
 */
export default function PanelSection(props) {
  return (
    <div
      className={css({
        backgroundColor: brandColors.reallyReallyLightGray,
        paddingBottom: '1rem'
      })}
    >
      <div
        className={cx(
          commonCss,
          css({
            borderBottom: `solid .1rem ${brandColors.reallyLightGray} !important`,
            paddingBottom: '1rem',
            paddingTop: '1rem'
          })
        )}
      >
        {props.title}
        {
          props.summary &&
          <div
            className={css({
              fontSize: '.8rem'
            })}
          >
            {props.summary}
          </div>
        }
      </div>
      {props.children}
    </div>

  )
}